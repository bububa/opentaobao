package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaTianjiSupplierOrderQuery 查询供应商订单
// alibaba.tianji.supplier.order.query
//
// 查询供应商订单
func AlibabaTianjiSupplierOrderQuery(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaTianjiSupplierOrderQueryAPIRequest, resp *alicom.AlibabaTianjiSupplierOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
