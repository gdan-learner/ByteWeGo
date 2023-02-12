// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	userservice "github.com/gdan0324/ByteWeGo/api/kitex_gen/userservice"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*userservice.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CheckUser":  kitex.NewMethodInfo(checkUserHandler, newUserServiceCheckUserArgs, newUserServiceCheckUserResult, false),
		"CreateUser": kitex.NewMethodInfo(createUserHandler, newUserServiceCreateUserArgs, newUserServiceCreateUserResult, false),
		"GetUser":    kitex.NewMethodInfo(getUserHandler, newUserServiceGetUserArgs, newUserServiceGetUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "userservice",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userservice.UserServiceCheckUserArgs)
	realResult := result.(*userservice.UserServiceCheckUserResult)
	success, err := handler.(userservice.UserService).CheckUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCheckUserArgs() interface{} {
	return userservice.NewUserServiceCheckUserArgs()
}

func newUserServiceCheckUserResult() interface{} {
	return userservice.NewUserServiceCheckUserResult()
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userservice.UserServiceCreateUserArgs)
	realResult := result.(*userservice.UserServiceCreateUserResult)
	success, err := handler.(userservice.UserService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return userservice.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return userservice.NewUserServiceCreateUserResult()
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userservice.UserServiceGetUserArgs)
	realResult := result.(*userservice.UserServiceGetUserResult)
	success, err := handler.(userservice.UserService).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserArgs() interface{} {
	return userservice.NewUserServiceGetUserArgs()
}

func newUserServiceGetUserResult() interface{} {
	return userservice.NewUserServiceGetUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CheckUser(ctx context.Context, req *userservice.CheckUserRequest) (r *userservice.CheckUserResponse, err error) {
	var _args userservice.UserServiceCheckUserArgs
	_args.Req = req
	var _result userservice.UserServiceCheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUser(ctx context.Context, req *userservice.CreateUserRequest) (r *userservice.CreateUserResponse, err error) {
	var _args userservice.UserServiceCreateUserArgs
	_args.Req = req
	var _result userservice.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, req *userservice.GetUserRequest) (r *userservice.GetUserResponse, err error) {
	var _args userservice.UserServiceGetUserArgs
	_args.Req = req
	var _result userservice.UserServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
