package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kuzja086/grpcTrain.git/pkg/apiClick"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	host     = ""
	UserName = ""
	Password = ""
)

func main() {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := apiClick.NewClickHouseClient(conn)
	queryInfo := &apiClick.QueryInfo{
		Query:    "Select top 3 * from httpExchange",
		Database: "log_storage_test",
		// OutputFormat: "json",
		UserName: UserName,
		Password: Password,
	}
	res, err := client.ExecuteQuery(context.Background(), queryInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res.GetOutput()))
}
