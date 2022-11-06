package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/service/user/model"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

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

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			logrus.Errorln(err.Error())
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &product.DetailResponse{
		Id:     res.Id,
		Name:   res.Name,
		Desc:   res.Desc,
		Stock:  res.Stock,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
