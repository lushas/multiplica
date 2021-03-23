package main

import (
        "fmt"
        "log"
        "net"
        "google.golang.org/grpc"
	pb "multiplica/multiplica"
	"context"
)

type Server struct {
	pb.UnimplementedMultiplicarServer
}


func main() {

        fmt.Println("Starting gRPC server")

        lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9500))
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }

        s := Server{}

        grpcServer := grpc.NewServer()

        pb.RegisterMultiplicarServer(grpcServer, &s)

        if err := grpcServer.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %s", err)
        }
}


func (s *Server) Multiplica(ctx context.Context, in *pb.MessageIn) (*pb.MessageOut, error) {
        log.Printf("Receive message body from client: %d, %d", in.Number1, in.Number2)
        result := in.Number1 * in.Number2
        return &pb.MessageOut{Result: result}, nil
}

