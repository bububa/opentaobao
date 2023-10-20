package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBcFutureStockOutbound 供应商期货出库
// alibaba.alihealth.bc.future.stock.outbound
//
// 供应商期货出库
func AlibabaAlihealthBcFutureStockOutbound(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBcFutureStockOutboundAPIRequest, resp *alihealth2.AlibabaAlihealthBcFutureStockOutboundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
