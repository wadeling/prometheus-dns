package handlers

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
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
	p := httputil.NewSingleHostReverseProxy(target)
	cp := x509.NewCertPool()
	content,err := ioutil.ReadFile("./ca.crt")
	if err != nil {
		return nil,err
	}
	ok := cp.AppendCertsFromPEM(content)
	if !ok {
		return nil,errors.New("append certs from pen error")
	}

	p.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: cp,
			//InsecureSkipVerify: true,
		},
	}
	return &ProxyService{
		proxy:  p,
		target: target,
	}, nil
}
