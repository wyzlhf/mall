package logic

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"mall/service/user/model"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"

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

func (l *DetailLogic) Detail(in *pay.DetailRequest) (*pay.DetailResponse, error) {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			logrus.Errorln(err.Error())
			return nil, status.Error(500, err.Error())
		}
		return nil, status.Error(500, err.Error())
	}

	return &pay.DetailResponse{
		Id:     res.Id,
		Uid:    res.Uid,
		Oid:    res.Oid,
		Amount: res.Amount,
		Source: res.Source,
		Status: res.Status,
	}, nil
}
