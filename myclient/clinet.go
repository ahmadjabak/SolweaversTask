package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    pb "github.com/ahmadjabak/SolweaversTask/example" 
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewMyServiceClient(conn)

    // Call your gRPC service method
    response, err := c.MyMethod(context.Background(), &pb.MyRequest{RequestData: "World"})
    if err != nil {
        log.Fatalf("Error calling MyMethod: %v", err)
    }

    log.Printf("Response from server: %s", response.ResponseData)
}
