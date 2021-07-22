package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/go-redis/redis"
	"google.golang.org/grpc"

	pb "github.com/tingyoulin/docker_goapp_test/pb/mypb"
)

var rdb = newClient() // declare redis client

/* Rest API
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(recordUserMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})

	router.GET("/ping1", func(c *gin.Context) {
		// member := c.Param("member")

		c.String(http.StatusOK, "Pong1")
	})

	router.GET("/get_visits", func(c *gin.Context) {
		// member := c.Param("member")
		userIP := c.ClientIP()
		val, err := rdb.Get(userIP).Result()

		switch {
		case err == redis.Nil:
			c.String(http.StatusBadRequest, fmt.Sprint("key does not exist"))
		case err != nil:
			c.String(http.StatusBadRequest, fmt.Sprint(err))
		case val == "":
			c.String(http.StatusBadRequest, fmt.Sprint("value is empty"))
		}

		c.String(http.StatusOK, fmt.Sprintf("Got!\n%s visits: %s", userIP, val))
	})

	router.GET("/delete/", func(c *gin.Context) {
		userIP := c.ClientIP()
		err := rdb.Del(userIP).Err()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprint(err))
		}
		c.String(http.StatusOK, fmt.Sprintf("%s deleted!", userIP))
	})
	return router
}

func recordUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIP := c.ClientIP()
		log.Printf("User IP: %s", userIP)

		err := rdb.Incr(userIP).Err()
		if err != nil {
			err := rdb.Set(userIP, 1, 0).Err()

			if err != nil {
				log.Println(fmt.Sprint(err))
			}

			log.Printf("create: %s\n", userIP)
			return
		}
	}
} */

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // redis 是 docker-compose.yml 裡的 service name (or alias)
		Password: "0000",
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatalln(err)
	}

	return client
}

// redis 操作
func incrName(name string) {
	err := rdb.Incr(name).Err()
	if err != nil {
		err := rdb.Set(name, 1, 0).Err()

		if err != nil {
			log.Println(err)
		}

		log.Printf("create: %s\n", name)
		return
	}
}

func getName(name string) (visits int) {
	val, err := rdb.Get(name).Result()

	switch {
	case err == redis.Nil:
		log.Println("key does not exist")
		return
	case err != nil:
		log.Println(err)
		return
	case val == "":
		log.Println("value is empty")
		return
	}

	visits, err = strconv.Atoi(val)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func deleteName(name string) (res string) {
	err := rdb.Del(name).Err()
	if err != nil {
		res = fmt.Sprint(err)
		return
	}
	res = fmt.Sprintf("%s deleted!", name)
	return
}

// type Member struct {
// 	name   string
// 	visits string
// }

type server struct{}

/* 實作 Protocol Buffer 中的 service */
func (*server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingReply, error) {
	name := req.GetName()
	incrName(name)
	res := &pb.PingReply{
		Result: name + " Pong",
	}
	return res, nil
}

func (*server) Ping1(ctx context.Context, req *pb.PingRequest) (*pb.PingReply, error) {
	name := req.GetName()
	incrName(name)
	res := &pb.PingReply{
		Result: name + " Pong1",
	}
	return res, nil
}

func (*server) GetVisits(ctx context.Context, req *pb.GetVisitsRequest) (*pb.GetVisitsReply, error) {
	name := req.GetName()

	res := &pb.GetVisitsReply{
		Visits: int32(getName(name)),
	}
	return res, nil
}

func (*server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	name := req.GetName()
	res := &pb.DeleteReply{
		Result: deleteName(name),
	}
	return res, nil
}

func main() {
	// router := setupRouter()
	// router.Run() // default localhost:50051

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGrpcTestServer(grpcServer, &server{})

	fmt.Println("gRPC Server Starting Listening at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
