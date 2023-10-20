package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBcFutureStockInbound 供应商上报期货库存
// alibaba.alihealth.bc.future.stock.inbound
//
// 供应商上报期货库存
func AlibabaAlihealthBcFutureStockInbound(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBcFutureStockInboundAPIRequest, session string) (*alihealth2.AlibabaAlihealthBcFutureStockInboundAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthBcFutureStockInboundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
