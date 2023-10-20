package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBcFutureStockInbound 供应商上报期货库存
// alibaba.alihealth.bc.future.stock.inbound
//
// 供应商上报期货库存
func AlibabaAlihealthBcFutureStockInbound(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBcFutureStockInboundAPIRequest, resp *alihealth2.AlibabaAlihealthBcFutureStockInboundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
