package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBcFutureStockOutbound 供应商期货出库
// alibaba.alihealth.bc.future.stock.outbound
//
// 供应商期货出库
func AlibabaAlihealthBcFutureStockOutbound(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBcFutureStockOutboundAPIRequest, resp *alihealth2.AlibabaAlihealthBcFutureStockOutboundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
