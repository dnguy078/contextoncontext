package main

import (
	"context"
	"fmt"
	"net/http"
)

func pingGoogle(ctx context.Context) error {
	res, err := http.Get("https://google.com")
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("got %s from google", res.StatusCode)
	}
	return nil
}
