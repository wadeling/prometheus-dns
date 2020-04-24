package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyService struct {
	target *url.URL
	proxy *httputil.ReverseProxy
}

func (p *ProxyService) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	p.proxy.ServeHTTP(rw, req)
}

func NewProxyService(backendServerAddr string) (*ProxyService, error) {
	var target, err = url.Parse(backendServerAddr)
	if err != nil {
		return nil, err
	}
	return &ProxyService{
		proxy:  httputil.NewSingleHostReverseProxy(target),
		target: target,
	}, nil
}
