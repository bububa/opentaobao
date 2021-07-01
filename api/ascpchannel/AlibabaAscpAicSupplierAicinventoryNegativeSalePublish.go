package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

/* AlibabaAscpAicSupplierAicinventoryNegativeSalePublish
AIC负卖库存新增和修改接口
alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish

新增负卖库存记录和变更负卖库存记录 */
func AlibabaAscpAicSupplierAicinventoryNegativeSalePublish(clt *core.SDKClient, req *ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest, session string) (*ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
