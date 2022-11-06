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

type PaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidLogic {
	return &PaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaidLogic) Paid(in *order.PaidRequest) (*order.PaidResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	res,err:=l.svcCtx.OrderModel.FindOne(l.ctx,in.Id)
	if err!=nil{
		if err==model.ErrNotFound{
			logrus.Errorln(err.Error())
			return nil, status.Error(100,"订单不存在")
		}
		return nil, status.Error(500,err.Error())
	}
	res.Status=1
	err=l.svcCtx.OrderModel.Update(l.ctx,res)
	if err!=nil {
		logrus.Errorln(err.Error())
		return nil,status.Error(500,err.Error())
	}

	return &order.PaidResponse{}, nil
}
