package reloader

import (
	"github.com/mduvall/fsnotify2"
	"log"
)

type Server struct {
	watcher *fsnotify2.Watcher
}

func CreateServer() (s *Server, err error) {
	watcher, err := fsnotify2.NewWatcher()

	if err != nil {
		log.Fatal("Unable to open directory")
		return nil, err
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println(ev)
				// publish every time this happens
			case err := <-watcher.Error:
				log.Println("error: ", err)
			}
		}
	}()

	err = watcher.WatchAllFlags("./", fsnotify2.FILE_WRITE)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Server{
		watcher: watcher,
	}, nil
}

func (s *Server) ListenAtPath(path string) {

}
