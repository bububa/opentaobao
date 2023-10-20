package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpaicsupplieraicinventorynegativesaleinvalidate 负卖库存失效接口
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate
//
// 失效负卖库存数据
func Alibabaascpaicsupplieraicinventorynegativesaleinvalidate(clt *core.SDKClient, req *ascpchannel.AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest, session string) (*ascpchannel.AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIResponse, error) {
	var resp ascpchannel.AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
