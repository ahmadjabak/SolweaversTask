package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/ahmadjabak/SolweaversTask/example" // Import the package generated from your .proto file
)

// Define your service implementation
type server struct{}

func (s *server) MyMethod(ctx context.Context, req *pb.MyRequest) (*pb.MyResponse, error) {
    // Implement the logic for your service's method
    response := &pb.MyResponse{ResponseData: "Hello, " + req.RequestData}
    return response, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterMyServiceServer(s, &server{})
    log.Println("Server is running on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
