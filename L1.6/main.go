package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func shutdownWithCtx() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина shutdownWithCtx работает")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
}

func shutdownWithNotify() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-exit:
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Горутина shutdownWithNotify работает")
		}

	}
}

func shutdownWithNotifyContext() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Горутина shutdownWithNotifyContext работает")
		}

	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Симулирование некоторой работы
	for i := 0; i < 3; i++ {
		fmt.Printf("Горутина shutdownWithGoexit %d обрабатывает задачу %d\n", id, i)
	}
	// Корректный выход из горутины
	runtime.Goexit()
}
func shutdownWithGoexit() {
	var wg sync.WaitGroup

	for i := 1; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}

func shutdownWithStruct() {
	stop := make(chan struct{}) // Канал для остановки

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("Горутина shutdownWithStruct работает")
			case <-stop:
				fmt.Println("Горутина завершает работу")
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(stop)
}

// L1.6 Остановка горутины
func main() {
	// Запускать по отдельности
	shutdownWithCtx()
	//shutdownWithNotify()
	//shutdownWithNotifyContext()
	//shutdownWithGoexit()
	//shutdownWithStruct()
}
