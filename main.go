package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("hello world")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond)
	defer cancel() // should always close the context else memory leak
	if err := pingGoogle(ctx); err != nil {
		log.Printf("could not ping google: %v", err)
		return
	}
	log.Println("wifi is up")
}
