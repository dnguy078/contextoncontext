package main

import (
	"context"
	"time"
)

func tree() {
	ctx := context.Background()
	ctx2, _ := context.WithCancel(ctx)
	ctx3, _ := context.WithTimeout(ctx2, time.Second*5)
	ctx4, _ := context.WithTimeout(ctx3, time.Second*4)
	ctx5, _ := context.WithValue(ctx3, "userID", 12)
}
