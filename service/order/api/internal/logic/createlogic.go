package logic

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
	"mall/service/order/rpc/types/order"
	"mall/service/product/rpc/types/product"
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

//	func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
//		//logrus.SetReportCaller(true)
//		//logrus.SetLevel(logrus.TraceLevel)
//		//res, err := l.svcCtx.OrderRpc.Create(l.ctx, &orderclient.CreateRequest{
//		//	Uid:    req.Uid,
//		//	Pid:    req.Pid,
//		//	Amount: req.Amount,
//		//	Status: req.Status,
//		//})
//		//if err != nil {
//		//	logrus.Errorln(err.Error())
//		//	return nil, err
//		//}
//		//
//		//return &types.CreateResponse{
//		//	Id: res.Id,
//		//}, nil
//		orderRpcBusiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
//		if err != nil {
//			return nil, status.Error(100, "订单创建异常")
//		}
//		productRpcBusiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
//		if err != nil {
//			return nil, status.Error(100, "订单创建异常")
//		}
//		var dtmServer = "etcd://127.0.0.1:2379/dtmservice"
//		gid := dtmgrpc.MustGenGid(dtmServer)
//		// 创建一个saga协议的事务
//		saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
//			Add(orderRpcBusiServer+"/order.Order/Create", orderRpcBusiServer+"/order.Order/CreateRevert", &order.CreateRequest{
//				Uid:    req.Uid,
//				Pid:    req.Pid,
//				Amount: req.Amount,
//				Status: 0,
//			}).
//			Add(productRpcBusiServer+"/product.Product/DecrStock", productRpcBusiServer+"/product.Product/DecrStockRevert", &product.DecrStockRequest{
//				Id:  req.Pid,
//				Num: 1,
//			})
//
//		// 事务提交
//		err = saga.Submit()
//		if err != nil {
//			return nil, status.Error(500, err.Error())
//		}
//
//		return &types.CreateResponse{}, nil
//	}
func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	logrus.SetLevel(logrus.WarnLevel)
	logrus.SetReportCaller(true)
	// 获取 OrderRpc BuildTarget
	orderRpcBusiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// 获取 ProductRpc BuildTarget
	productRpcBusiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// dtm 服务的 etcd 注册地址
	var dtmServer = "etcd://127.0.0.1:2379/dtmservice"
	// 创建一个gid
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
	logrus.Error("错误出自此处", err.Error())
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateResponse{}, nil
}
