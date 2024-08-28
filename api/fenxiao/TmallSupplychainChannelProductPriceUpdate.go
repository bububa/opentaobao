package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductPriceUpdate 渠道价格更新接口
// tmall.supplychain.channel.product.price.update
//
// 更新渠道产品价格
func TmallSupplychainChannelProductPriceUpdate(ctx context.Context, clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductPriceUpdateAPIRequest, resp *fenxiao.TmallSupplychainChannelProductPriceUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
