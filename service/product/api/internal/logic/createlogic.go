package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"mall/service/product/rpc/productclient"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"

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

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	res, err := l.svcCtx.ProductRpc.Create(l.ctx, &productclient.CreateRequest{
		Name:   req.Name,
		Desc:   req.Desc,
		Stock:  req.Stock,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}
	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
