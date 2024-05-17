package main

import (
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

func get_client() http.Client {

	proxy, err := proxy.SOCKS5("tcp", app.Proxy.Socks5, nil, proxy.Direct)

	if err != nil {
		return http.Client{Timeout: time.Second * 60}
	}
	return http.Client{
		Transport: &http.Transport{
			Dial: proxy.Dial,
		},

		Timeout: time.Second * 60,
	}
}
