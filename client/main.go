package main

import (
	"log"
	"net/http"

	pb "github.com/ronexlemon/godocker/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddress = "localhost:9001"
	httpAddr = ":8080"
)

func main(){
	conn,err := grpc.NewClient(serverAddress,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to listen server at : %s",serverAddress)
		}
		mux := http.NewServeMux()
		client := pb.NewStoreServiceClient(conn)
		
		handler := NewClientHandler(client)
		handler.registerRoute(mux)
		log.Println("Server is listening at : ",serverAddress)
		if err := http.ListenAndServe(httpAddr, mux); err != nil {
			log.Fatal("failed to start server")
		}

}