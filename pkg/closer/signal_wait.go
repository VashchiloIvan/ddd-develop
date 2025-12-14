package closer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var expectedSignals = []os.Signal{
	syscall.SIGINT, syscall.SIGTERM,
}

func WaitInterruptSignal(cancel context.CancelFunc) error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, expectedSignals...)

	fmt.Println("Received termination signal, shutting down...") // todo to tog.Onfo

	<-signalChan

	cancel()

	return globalCloser.closeAll()
}
