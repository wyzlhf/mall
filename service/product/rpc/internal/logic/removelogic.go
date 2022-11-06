package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/service/user/model"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *product.RemoveRequest) (*product.RemoveResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			logrus.Errorln(err.Error())
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	err = l.svcCtx.ProductModel.Delete(l.ctx, res.Id)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, status.Error(500, err.Error())
	}
	return &product.RemoveResponse{}, nil
}
