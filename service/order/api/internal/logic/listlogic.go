package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"mall/service/order/rpc/orderclient"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (resp *types.OrdersResponse, err error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	res,err:=l.svcCtx.OrderRpc.List(l.ctx,&orderclient.ListRequest{
		Uid: req.Uid,
	})
	if err!=nil{
		logrus.Errorln(err.Error())
		return nil,err
	}
	orderList:=make([]*types.OrderResponse,0)
	for _,item:=range res.Data{
		orderList=append(orderList,&types.OrderResponse{
			Id: item.Id,
			Uid: item.Uid,
			Pid: item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}
	ordersResponse:=&types.OrdersResponse{
		Orders: orderList,
	}

	return ordersResponse,nil
}
