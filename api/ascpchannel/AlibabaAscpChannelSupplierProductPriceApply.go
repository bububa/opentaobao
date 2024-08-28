package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSupplierProductPriceApply 供应商设置渠道产品价格
// alibaba.ascp.channel.supplier.product.price.apply
//
// 供应商设置渠道产品价格
func AlibabaAscpChannelSupplierProductPriceApply(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSupplierProductPriceApplyAPIRequest, resp *ascpchannel.AlibabaAscpChannelSupplierProductPriceApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
