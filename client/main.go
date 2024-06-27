package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "56HW/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:9779", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getProductReq := &pb.ProductRequest{Id: "123"}
	getProductRes, err := client.GetProduct(ctx, getProductReq)
	if err != nil {
		log.Fatalf("Error while calling GetProduct: %v", err)
	}
	log.Printf("GetProduct response: %v", getProductRes)

	addProductReq := &pb.AddProductRequest{
		Id:    "456",
		Name:  "New Product",
		Price: 199.99,
	}
	addProductRes, err := client.AddProduct(ctx, addProductReq)
	if err != nil {
		log.Fatalf("Error while calling AddProduct: %v", err)
	}
	log.Printf("AddProduct response: %v", addProductRes)

	deleteProductReq := &pb.ProductRequest{Id: "456"}
	deleteProductRes, err := client.DeleteProduct(ctx, deleteProductReq)
	if err != nil {
		log.Fatalf("Error while calling DeleteProduct: %v", err)
	}
	log.Printf("DeleteProduct response: %v", deleteProductRes)
}
