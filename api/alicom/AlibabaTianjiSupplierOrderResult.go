package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaTianjiSupplierOrderResult 供应商处理订单接口（订购成功/失败、发货）
// alibaba.tianji.supplier.order.result
//
// 供应商处理订单接口（订购成功/失败、发货）
func AlibabaTianjiSupplierOrderResult(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaTianjiSupplierOrderResultAPIRequest, resp *alicom.AlibabaTianjiSupplierOrderResultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
