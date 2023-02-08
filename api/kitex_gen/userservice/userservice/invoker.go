// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	server "github.com/cloudwego/kitex/server"
	userservice "github.com/gdan0324/ByteWeGo/api/kitex_gen/userservice"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler userservice.UserService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
