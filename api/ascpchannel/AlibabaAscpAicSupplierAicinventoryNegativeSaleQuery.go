package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpaicsupplieraicinventorynegativesalequery 商家查询负卖库存
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.query
//
// 商家根据当前接口查询负卖货品的库存
func Alibabaascpaicsupplieraicinventorynegativesalequery(clt *core.SDKClient, req *ascpchannel.AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest, session string) (*ascpchannel.AlibabaascpaicsupplieraicinventorynegativesalequeryAPIResponse, error) {
	var resp ascpchannel.AlibabaascpaicsupplieraicinventorynegativesalequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
