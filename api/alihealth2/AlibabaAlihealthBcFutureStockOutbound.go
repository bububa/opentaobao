package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBcFutureStockOutbound 供应商期货出库
// alibaba.alihealth.bc.future.stock.outbound
//
// 供应商期货出库
func AlibabaAlihealthBcFutureStockOutbound(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBcFutureStockOutboundAPIRequest, session string) (*alihealth2.AlibabaAlihealthBcFutureStockOutboundAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthBcFutureStockOutboundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
