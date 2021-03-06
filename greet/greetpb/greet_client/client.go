package main

import (
	"fmt"
	"log"

	"github.com/Tambarie/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)




func main()  {
	fmt.Println("connected to client")
	cc, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connect: %v",err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("created client: %f",c)
}