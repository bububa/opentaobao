package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidate 负卖库存失效接口
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate
//
// 失效负卖库存数据
func AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidate(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest, resp *ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
