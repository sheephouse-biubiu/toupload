package main

import (
	"context"
	"doupload/uploadservice"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	fmt.Println("esay to upload, easy to code!!!")
	ctx := context.Background()
	r := weaver.Init(ctx)
	uploadService, err := weaver.Get[uploadservice.T](r)
	if err != nil {
		log.Fatal(err)
	}
	opts := weaver.ListenerOptions{LocalAddress: "0.0.0.0:10028"}
	lis, err := r.Listener("upload", opts)
	if err != nil {
		log.Fatal(err)
	}
	uploadService.RegisterRoutes(ctx)
	http.Serve(lis, nil)
}
