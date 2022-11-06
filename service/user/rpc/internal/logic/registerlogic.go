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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == nil {
		logrus.Errorln("该用户已存在")
		return nil, status.Error(100, "该用户已存在")
	}
	if err == model.ErrNotFound {
		newUser := model.User{
			Name:     in.Name,
			Gender:   in.Gender,
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			logrus.Errorln(err.Error())
			return nil, status.Error(500, err.Error())
		}
		userId, err := res.LastInsertId()
		if err != nil {
			logrus.Errorln(err.Error())
			return nil, status.Error(500, err.Error())
		}
		//newUser.Id=uint64(userId)
		return &user.RegisterResponse{
			Id:     uint64(userId),
			Name:   newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		}, nil
	}
	//logrus.Errorln(err.Error())
	return nil, status.Error(500, err.Error())
}
