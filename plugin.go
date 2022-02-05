package main

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"fmt"
)

func newPlugin() (*plugin, error) {
	return &plugin{}, nil
}

type plugin struct {
}

func (p *plugin) AuthZReq(req authorization.Request) authorization.Response {
	fmt.Printf("AuthZReq user: %s\n", req.User)
	// reqURI, _ := url.QueryUnescape(req.RequestURI)
	// reqURL, _ := url.ParseRequestURI(reqURI)

	// obj := reqURL.String()
	// act := req.RequestMethod
	// fmt.Printf("obj: %s - act: %s\n",obj,act)
	if req.User == "docker_client" {
		return authorization.Response{Allow: true}
	} else {
	  return authorization.Response{Allow: false}
	}
}

func (p *plugin) AuthZRes(req authorization.Request) authorization.Response {
	fmt.Printf("AuthZReq user: %s\n", req.User)
	return authorization.Response{Allow: true}
}
