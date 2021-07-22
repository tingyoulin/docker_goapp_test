package main

import (
	"context"
	"log"

	pb "github.com/tingyoulin/docker_goapp_test/pb/mypb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewGrpcTestClient(conn)

	doUnary(client)
}

func doUnary(client pb.GrpcTestClient) {
	pingReq := &pb.PingRequest{
		Name: "Rayne",
	}

	for i := 0; i < 10; i++ {
		var (
			res *pb.PingReply
			err error
		)

		if i%3 == 0 {
			res, err = client.Ping(context.Background(), pingReq)
		} else {
			res, err = client.Ping1(context.Background(), pingReq)
		}

		if err != nil {
			log.Fatalf("error while calling GrpcTestService: %v \n", err)
		}
		log.Printf("Response from GrpcTestService: %v", res.Result)
	}
	getReq := &pb.GetVisitsRequest{
		Name: "Rayne",
	}
	getRes1, err := client.GetVisits(context.Background(), getReq)
	if err != nil {
		log.Fatalf("error while calling GrpcTestService: %v \n", err)
	}
	log.Printf("Response from GrpcTestService: visits: %v times", getRes1.Visits)

	delReq := &pb.DeleteRequest{
		Name: "Rayne",
	}
	delRes, err := client.Delete(context.Background(), delReq)
	if err != nil {
		log.Fatalf("error while calling GrpcTestService: %v \n", err)
	}
	log.Printf("Response from GrpcTestService: %v", delRes.Result)

	getRes2, err := client.GetVisits(context.Background(), getReq)
	if err != nil {
		log.Fatalf("error while calling GrpcTestService: %v \n", err)
	}
	log.Printf("Response from GrpcTestService: visits: %v times", getRes2.Visits)
}
