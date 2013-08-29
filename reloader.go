package main

import (
	"github.com/mduvall/fsnotify2"
	"github.com/mduvall/reloader/reloader"
	"log"
)

func main() {
	watcher, err := fsnotify2.NewWatcher()

	if err != nil {
		log.Fatal("Unable to open directory")
	}

	done := make(chan bool)
	transport := createTransport()

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				transport.Publish(ev.Name)
			case err := <-watcher.Error:
				log.Println("error: ", err)
			}
		}
	}()

	err = watcher.WatchAllFlags("./", fsnotify2.FILE_WRITE)

	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}

func createTransport() *reloader.Transport {
	return &reloader.Transport{}
}
