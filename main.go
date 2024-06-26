package main

import (
	"context"
	"os"
	"os/signal"
	"xploit/cmd"
    "syscall"
)

func main() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

    for {
        ctx, cancel := context.WithCancel(context.Background())
        go func() {
            <-sigChan
            cancel()
        }()

        cmd.InteractiveShell(ctx)
    }
}
