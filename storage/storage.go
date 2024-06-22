package storage

import (
	"context"
	op "uzum_orderclone/genproto/order_product"
	st "uzum_orderclone/genproto/order_status_notes"
	ct "uzum_orderclone/genproto/orders_service"
)

type StorageI interface {
	CloseDB()
	Order() OrderyRepoI
	OrderProduct() OrderProductepoI
	OrderStatus() OrderStatusI
}

type OrderyRepoI interface {
	Create(ctx context.Context, req *ct.CreateOrder) (resp *ct.GetOrder, err error)
	GetByID(ctx context.Context, req *ct.OrderPrimaryKey) (resp *ct.GetOrder, err error)
	//GetALL(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error)
	//Update(ctx context.Context, req *ct.UpdateOrder) (resp *ct.GetOrder, err error)
	//Delete(ctx context.Context, req *ct.OrderPrimaryKey) (resp *ct.Empty, err error)
}

type OrderProductepoI interface {
	Create(ctx context.Context, req *op.CreateOrderProducts) (resp *op.GetOrderProducts, err error)
	GetByID(ctx context.Context, req *op.OrderProductsPrimaryKey) (resp *op.GetOrderProducts, err error)
	//GetALL(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error)
	//Update(ctx context.Context, req *ct.UpdateOrder) (resp *ct.GetOrder, err error)
	//Delete(ctx context.Context, req *ct.OrderPrimaryKey) (resp *ct.Empty, err error)
}

type OrderStatusI interface {
	Create(ctx context.Context, req *st.CreateStatusRequest) (resp *st.GetOrderStatusResponse, err error)
	GetByID(ctx context.Context, req *st.OrderPrimaryKeyRequest) (resp *st.GetOrderStatusResponse, err error)
	PUTCH(ctx context.Context, req *st.OrderPrimaryStatusKeyRequest) (resp *st.GetOrderStatusResponse, err error)
	GetStatusByID(ctx context.Context, req *st.OrderPrimaryStatusKeyRequest) (resp *st.GetOrderStatusResponse, err error)
	// //GetALL(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error)
	// //Update(ctx context.Context, req *ct.UpdateOrder) (resp *ct.GetOrder, err error)
	// //Delete(ctx context.Context, req *ct.OrderPrimaryKey) (resp *ct.Empty, err error)
}
