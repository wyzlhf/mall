package logic

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/types/order"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
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
	//logrus.SetReportCaller(true)
	//logrus.SetLevel(logrus.TraceLevel)
	//res, err := l.svcCtx.OrderRpc.Create(l.ctx, &orderclient.CreateRequest{
	//	Uid:    req.Uid,
	//	Pid:    req.Pid,
	//	Amount: req.Amount,
	//	Status: req.Status,
	//})
	//if err != nil {
	//	logrus.Errorln(err.Error())
	//	return nil, err
	//}
	//
	//return &types.CreateResponse{
	//	Id: res.Id,
	//}, nil
	orderRpcBusiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}
	productRpcBusiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}
	var dtmServer = "etcd://127.0.0.1:2379/dtmservice"
	gid := dtmgrpc.MustGenGid(dtmServer)
	// 创建一个saga协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcBusiServer+"/order.Order/Create", orderRpcBusiServer+"/order.Order/CreateRevert", &order.CreateRequest{
			Uid:    req.Uid,
			Pid:    req.Pid,
			Amount: req.Amount,
			Status: 0,
		}).
		Add(productRpcBusiServer+"/product.Product/DecrStock", productRpcBusiServer+"/product.Product/DecrStockRevert", &product.DecrStockRequest{
			Id:  req.Pid,
			Num: 1,
		})

	// 事务提交
	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateResponse{}, nil
}
