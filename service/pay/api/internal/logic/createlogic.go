package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"mall/service/pay/rpc/payclient"

	"mall/service/pay/api/internal/svc"
	"mall/service/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateReuest) (resp *types.CreateResponse, err error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	res,err:=l.svcCtx.PayRpc.Create(l.ctx,&payclient.CreateRequest{
		Uid: req.Uid,
		Oid: req.Oid,
		Amount: req.Amount,
	})
	if err!=nil{
		logrus.Error(err.Error())
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	},nil
}
