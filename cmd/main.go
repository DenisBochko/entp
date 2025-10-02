package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DenisBochko/entp"
)

func main() {
	ctx := context.Background()

	client := entp.NewClient()

	t, err := client.Now(ctx)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}

	println(t.Format(time.RFC3339Nano))
}
