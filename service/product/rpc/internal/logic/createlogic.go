package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/service/product/model"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

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

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	newProduct := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
		Status: in.Status,
	}
	res, err := l.svcCtx.ProductModel.Insert(l.ctx, &newProduct)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, status.Error(500, err.Error())
	}
	newProductId, err := res.LastInsertId()
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, status.Error(500, err.Error())
	}
	return &product.CreateResponse{
		Id: uint64(newProductId),
	}, nil
}
