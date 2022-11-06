package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/common/cryptx"
	"mall/service/user/model"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	res,err:=l.svcCtx.UserModel.FindOneByMobile(l.ctx,in.Mobile)
	if err!=nil{
		if err==model.ErrNotFound{
			logrus.Errorln("用户不存在")
			return nil,status.Error(100,"用户不存在")
		}
		logrus.Errorln(err.Error())
		return nil, status.Error(500,err.Error())
	}
	password:=cryptx.PasswordEncrypt(l.svcCtx.Config.Salt,in.Password)
	if password!=res.Password{
		logrus.Errorln("密码错误")
		return nil,status.Error(100,"密码错误")
	}
	return &user.LoginResponse{
		Id: res.Id,
		Name: res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	},nil
}
