package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/service/order/model"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *order.DetailRequest) (*order.DetailResponse, error) {
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

	return &order.DetailResponse{
		Id: res.Id,
		Uid: res.Uid,
		Pid: res.Pid,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
