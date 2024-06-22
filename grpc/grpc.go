package grpc

import (
	"uzum_orderclone/config"
	"uzum_orderclone/genproto/order_product"
	"uzum_orderclone/genproto/order_status_notes"
	"uzum_orderclone/genproto/orders_service"
	"uzum_orderclone/grpc/client"
	"uzum_orderclone/grpc/service"
	"uzum_orderclone/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	orders_service.RegisterOrderServiceServer(grpcServer, service.NewOrderService(cfg, log, strg, srvc))
	order_product.RegisterOrderProductsServiceServer(grpcServer, service.NewProductOrderService(cfg, log, strg, srvc))
	order_status_notes.RegisterOrderStatusServiceServer(grpcServer, service.NewOrderstatus(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
