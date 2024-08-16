package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"gadget-crafted/internal/config"
	"gadget-crafted/internal/handler"
	"gadget-crafted/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gadget_crafted.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	if !ctx.Init() {
		os.Exit(2)
	}

	var server *rest.Server
	if c.Cors == 1 {
		logx.Info("enable cors")
		server = rest.MustNewServer(c.RestConf, rest.WithCors())
	} else {
		logx.Info("disable cores")
		server = rest.MustNewServer(c.RestConf)
	}

	defer server.Stop()
	handler.RegisterHandlers(server, ctx)

	logx.Infof("Starting gadget_crafted server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
