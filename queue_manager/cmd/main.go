package main

import (
	"context"
	"log"

	cmd "github.com/DimTur/lp_queue_manager/cmd/queue_manager"
)

func main() {
	ctx := context.Background()

	cmd := cmd.NewInitCmd()
	if err := cmd.ExecuteContext(ctx); err != nil {
		log.Fatalf("smth went wrong: %s", err)
	}
}
