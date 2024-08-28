package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSupplierProductGoodsBind 渠道产品与货品绑定接口
// alibaba.ascp.channel.supplier.product.goods.bind
//
// 渠道产品与货品绑定接口
func AlibabaAscpChannelSupplierProductGoodsBind(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSupplierProductGoodsBindAPIRequest, resp *ascpchannel.AlibabaAscpChannelSupplierProductGoodsBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
