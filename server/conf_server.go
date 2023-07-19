package server

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/zjmnssy/ezlog/impl"
	"sync"
)

const (
	defaultPipeFile = "./changeLogConfigPipe"
)

type ConfigServer struct {
	pipeServer *PipeServer
}

func newConfigServer() *ConfigServer {
	ps := NewPipeServer(defaultPipeFile, 0)

	return &ConfigServer{pipeServer: ps}
}

func (cs *ConfigServer) Start() {
	err := cs.pipeServer.Start()
	if err != nil {
		return
	}

	ch := cs.pipeServer.GetDataCh()

	go func() {
		for {
			data, ok := <-ch
			if ok {
				var json = jsoniter.ConfigCompatibleWithStandardLibrary
				err = json.Unmarshal([]byte(data), &impl.GetManager().Conf)
				if err == nil {
					impl.GetManager().SetConfig(impl.GetManager().Conf)
				} else {
					fmt.Printf("[WARN] data = %s, json unmarshal error = %s\n", data, err)
				}
			} else {
				fmt.Printf("data channel closed\n")
				return
			}
		}
	}()
}

func (cs *ConfigServer) Close() {
	cs.pipeServer.Close()
}

var once sync.Once
var instance *ConfigServer

func GetInstance() *ConfigServer {
	once.Do(func() {
		instance = newConfigServer()
	})

	return instance
}
