package main

import (
	"context"
	"time"
)

func PleaseCancelMe(parentCtx context.Context) {
	ctx, cancel := context.WithTimeout(parentCtx, time.Second)
	// Will not garbage collect before timer expires
	defer cancel()
}
