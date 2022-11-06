package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/orderclient"
	"mall/service/user/model"
	"mall/service/user/rpc/userclient"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	_,err:=l.svcCtx.UserRpc.UserInfo(l.ctx,&userclient.UserInfoRequest{
		Id: in.Uid,
	})
	if err!=nil{
		logrus.Error(err.Error())
		return nil, err
	}
	_,err=l.svcCtx.OrderRpc.Detail(l.ctx,&orderclient.DetailRequest{
		Id: in.Oid,
	})
	if err!=nil{
		logrus.Error(err.Error())
		return nil, err
	}
	res,err:=l.svcCtx.PayModel.FindOne(l.ctx,in.Id)
	if err!=nil{
		if err==model.ErrNotFound{
			logrus.Errorln(err.Error())
			return nil,status.Error(100,"支付不存在")
		}
		logrus.Errorln(err.Error())
		return nil, status.Error(500,err.Error())
	}
	if in.Amount!=res.Amount{
		logrus.Errorln(err.Error())
		return nil,status.Error(100,"支付金额与订单金额不符")
	}
	res.Source=in.Source
	res.Status=in.Status

	err=l.svcCtx.PayModel.Update(l.ctx,res)
	if err!=nil{
		return nil, status.Error(500,err.Error())
	}

	_,err=l.svcCtx.OrderRpc.Paid(l.ctx,&orderclient.PaidRequest{
		Id: in.Oid,
	})
	if err!=nil{
		return nil, status.Error(500,err.Error())
	}

	return &pay.CallbackResponse{}, nil
}
