package reloader

import (
	"github.com/mduvall/fsnotify2"
	"log"
	"net"
	"net/rpc"
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

type DummyEventEmitter struct{}
type Response struct {
	Sup string
}
type Request struct {
	Sup string
}

func (d *DummyEventEmitter) Sup(request Request, res *Response) error {
	res.Sup = "SUP!"
	return nil
}

func (s *Server) Start(path string) error {
	rpc.Register(&DummyEventEmitter{})
	listener, err := net.Listen("unix", path)

	if err != nil {
		return err
	}

	go rpc.Accept(listener)

	return nil
}
