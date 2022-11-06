package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/service/user/model"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

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

func (l *RemoveLogic) Remove(in *order.RemoveRequest) (*order.RemoveResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	res,err:=l.svcCtx.OrderModel.FindOne(l.ctx,in.Id)
	if err!=nil {
		if err==model.ErrNotFound{
			logrus.Errorln(err.Error())
			return nil, status.Error(100,"订单不存在")
		}
		return nil, status.Error(500,err.Error())
	}
	err=l.svcCtx.OrderModel.Delete(l.ctx,res.Id)
	if err!=nil{
		logrus.Errorln(err.Error())
		return nil, status.Error(500,err.Error())
	}

	return &order.RemoveResponse{}, nil
}
