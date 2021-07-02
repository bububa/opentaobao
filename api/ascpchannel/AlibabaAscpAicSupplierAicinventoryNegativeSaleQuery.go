package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpAicSupplierAicinventoryNegativeSaleQuery 商家查询负卖库存
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.query
//
// 商家根据当前接口查询负卖货品的库存
func AlibabaAscpAicSupplierAicinventoryNegativeSaleQuery(clt *core.SDKClient, req *ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest, session string) (*ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
