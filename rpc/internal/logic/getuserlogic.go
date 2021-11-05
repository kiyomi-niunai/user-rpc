package logic

import (
	"context"
	"fmt"
	"github.com/kiyomi-niunai/user/blob/master/rpc/model"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/kiyomi-niunai/user/blob/master/rpc/internal/svc"
	"github.com/kiyomi-niunai/user/blob/master/rpc/user"
	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line]
	var conn sqlx.SqlConn
	var c cache.CacheConf
	userObj, err := model.NewUsersModel(conn, c).FindOne(10000619)
	if err != nil {
		fmt.Println("报错的是", err)
	}
	return &user.UserResponse{
		Id: string(userObj.Id),
		Name: userObj.Name,
	}, nil
}
