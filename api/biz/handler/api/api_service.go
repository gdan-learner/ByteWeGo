// Code generated by hertz generator.

package api

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "github.com/gdan0324/ByteWeGo/api/biz/model/api"
	"github.com/gdan0324/ByteWeGo/api/biz/rpc"
	"github.com/gdan0324/ByteWeGo/api/kitex_gen/userservice"
)

// CheckUser .
// @router /douyin/user/login [POST]
func CheckUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CheckUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &userservice.CheckUserResponse{
			StatusCode: 1,
			StatusMsg:  "缺少用户名或密码",
		})
		return
	}

	log.Println(req)
	resp, err := rpc.CheckUser(context.Background(), &userservice.CheckUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// CreateUser .
// @router /douyin/user/register [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &userservice.CheckUserResponse{
			StatusCode: 1,
			StatusMsg:  "缺少用户名或密码",
		})
		return
	}

	resp, err := rpc.CreateUser(context.Background(), &userservice.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetUser .
// @router /douyin/user [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &userservice.CheckUserResponse{
			StatusCode: 1,
		})
		return
	}

	resp, err := rpc.GetUser(context.Background(), &userservice.GetUserRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}
