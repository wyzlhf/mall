package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/orderclient"
	"mall/service/pay/model"
	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pay.CreateRequest) (*pay.CreateResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &orderclient.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}
	_, err = l.svcCtx.PayModel.FindOneByOid(l.ctx, in.Oid)
	if err == nil {
		logrus.Errorln(err.Error())
		return nil, status.Error(100, "订单已创建")
	}
	newPay := model.Pay{
		Uid:    in.Uid,
		Oid:    in.Oid,
		Amount: in.Amount,
		Source: 0,
		Status: 0,
	}
	res, err := l.svcCtx.PayModel.Insert(l.ctx, &newPay)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newPayId, err := res.LastInsertId()
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, status.Error(500, err.Error())
	}
	return &pay.CreateResponse{
		Id: uint64(newPayId),
	}, nil
}
