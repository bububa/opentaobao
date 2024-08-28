package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductPriceGet 渠道价格查询接口
// tmall.supplychain.channel.product.price.get
//
// 渠道价格查询接口
func TmallSupplychainChannelProductPriceGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductPriceGetAPIRequest, resp *fenxiao.TmallSupplychainChannelProductPriceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
