package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    pb "56HW/proto"
)

type ProductServiceServer struct {
    pb.UnimplementedProductServiceServer
}

func (s *ProductServiceServer) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
    return &pb.ProductResponse{
        Id:    req.GetId(),
        Name:  "Sample Product",
        Price: 99.99,
    }, nil
}

func (s *ProductServiceServer) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.ProductResponse, error) {
    return &pb.ProductResponse{
        Id:    req.GetId(),
        Name:  req.GetName(),
        Price: req.GetPrice(),
    }, nil
}

func (s *ProductServiceServer) DeleteProduct(ctx context.Context, req *pb.ProductRequest) (*pb.DeleteProductResponse, error) {
    return &pb.DeleteProductResponse{
        Success: true,
    }, nil
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (interface{}, error) {
        log.Printf("Received request: %v", info.FullMethod)
        res, err := handler(ctx, req)
        log.Printf("Sending response: %v, error: %v", res, err)
        return res, err
    }
}

func main() {
    lis, err := net.Listen("tcp", ":9779")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor()))

    pb.RegisterProductServiceServer(s, &ProductServiceServer{})
    reflection.Register(s)

    log.Println("Server is running on port :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
