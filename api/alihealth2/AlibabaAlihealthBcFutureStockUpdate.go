package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBcFutureStockUpdate 供应商更新期货库存
// alibaba.alihealth.bc.future.stock.update
//
// 供应商更新期货库存
func AlibabaAlihealthBcFutureStockUpdate(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBcFutureStockUpdateAPIRequest, session string) (*alihealth2.AlibabaAlihealthBcFutureStockUpdateAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthBcFutureStockUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
