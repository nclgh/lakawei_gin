package lakawei_gin

import (
	"os"
	"fmt"
	"syscall"
	"os/signal"
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/nclgh/lakawei_discover"
)

type Gin struct {
	Engine *gin.Engine
}

var (
	HttpServer *Gin
)

func Init() *Gin {
	// ENV > FILE
	initConfigFromFile()
	initConfigFromENV()

	// 加载完配置后检查配置
	checkConfig()
	HttpServer = &Gin{
		Engine: gin.Default(),
	}
	return HttpServer
}

func Run() error {
	lakawei_discover.Register(ServiceName, fmt.Sprintf("%s:%s", ServiceAddr, ServicePort))
	errCh := make(chan error, 1)
	go func() {
		errCh <- HttpServer.Engine.Run(":" + ServicePort)
	}()
	return waitSignal(errCh)
}

func waitSignal(errCh <-chan error) error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	defer lakawei_discover.Unregister()
	for {
		select {
		case sig := <-ch:
			fmt.Printf("Got signal: %s, Exit..\n", sig)
			return errors.New(sig.String())
		case err := <-errCh:
			fmt.Printf("Engine run error: %s, Exit..\n", err)
			return err
		}
	}
}
