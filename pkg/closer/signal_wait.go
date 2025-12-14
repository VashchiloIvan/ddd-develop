package closer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var expectedSignals []syscall.Signal{
	syscall.SIGINT, syscall.SIGTERM
}

func WaitContextSignal(cancel context.CancelFunc) error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Received termination signal, shutting down...")

	<-signalChan

	cancel()

	return globalCloser.closeAll()
}
