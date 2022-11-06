package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"mall/service/product/rpc/productclient"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	_, err = l.svcCtx.ProductRpc.Remove(l.ctx, &productclient.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		logrus.Errorln(err.Error())
	}

	return &types.RemoveResponse{}, nil
}
