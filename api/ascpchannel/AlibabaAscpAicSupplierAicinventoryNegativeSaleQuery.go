package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpAicSupplierAicinventoryNegativeSaleQuery 商家查询负卖库存
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.query
//
// 商家根据当前接口查询负卖货品的库存
func AlibabaAscpAicSupplierAicinventoryNegativeSaleQuery(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest, resp *ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
