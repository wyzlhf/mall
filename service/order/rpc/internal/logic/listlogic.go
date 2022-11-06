package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/service/user/model"
	"mall/service/user/rpc/userclient"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListRequest) (*order.ListResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	_,err:=l.svcCtx.UserRpc.UserInfo(l.ctx,&userclient.UserInfoRequest{
		Id: in.Uid,
	})
	if err!=nil{
		logrus.Errorln(err.Error())
		return nil, err
	}
	list,err:=l.svcCtx.OrderModel.FindAllByUid(l.ctx,in.Uid)
	if err!=nil{
		if err==model.ErrNotFound{
			logrus.Errorln(err.Error())
			return nil,status.Error(100,"订单不存在")
		}
		logrus.Errorln(err.Error())
		return nil, status.Error(500,err.Error())
	}
	orderList:=make([]*order.DetailResponse,0)
	for _,item:=range list{
		orderList=append(orderList,&order.DetailResponse{
			Id: item.Id,
			Uid: item.Uid,
			Pid: item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}
	return &order.ListResponse{
		Data: orderList,
	}, nil
}
