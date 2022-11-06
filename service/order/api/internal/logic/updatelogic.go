package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"mall/service/order/rpc/orderclient"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	_,err=l.svcCtx.OrderRpc.Update(l.ctx,&orderclient.UpdateRequest{
		Id: req.Id,
		Uid: req.Uid,
		Pid: req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err!=nil{
		logrus.Errorln(err.Error())
		return nil,err
	}

	return &types.UpdateResponse{},nil
}
