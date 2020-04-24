package main

import (
	"fmt"
	"github.com/wadeling/prometheus-proxy/pkg/handlers"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
	//"go.uber.org/zap"

	//"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//"github.com/prometheus/common/model"
	//"github.com/prometheus/prometheus/prompb"
	"github.com/wadeling/prometheus-proxy/pkg/logger"
	//"io/ioutil"
	"net/http"
	"os"
	//"sync"
)

type prometheusProxy struct {
	listenPort string
	listenIP string
	backendServerAddr string
}

var log = logger.NewLogger(true)

func defaultProxy() *prometheusProxy{
	return &prometheusProxy{
		listenPort: "9100",
	}
}

func GetCmd() *cobra.Command {
	p := defaultProxy()

	cmd := &cobra.Command{
		Use:   "prometheus-proxy",
		Short: "prometheus-proxy",
		Run: func(cmd *cobra.Command, args []string) {
			p.run()
		},
	}

	f := cmd.PersistentFlags()
	f.StringVarP(&p.listenPort, "port", "p", p.listenPort, "TCP port to use for proxy")
	f.StringVarP(&p.listenIP, "listen IP", "i", p.listenIP, "TCP IP to use for proxy")
	f.StringVarP(&p.backendServerAddr, "backend Server Addr", "b", p.backendServerAddr, "backend Server Addr")

	return cmd
}

func (p *prometheusProxy) run() error {
	log.Info("run")

	addr := fmt.Sprintf("%s:%s",p.listenIP,p.listenPort)

	ps, err := handlers.NewProxyService(p.backendServerAddr)
	if err != nil {
		log.Error("new proxy service err",zap.Any("err",err))
		return err
	}
	return http.ListenAndServe(addr, ps)
}

func main() {
	//log.Info("hello proxy")

	cmd := GetCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}