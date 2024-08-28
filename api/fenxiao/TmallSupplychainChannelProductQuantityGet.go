package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductQuantityGet 渠道库存查询接口
// tmall.supplychain.channel.product.quantity.get
//
// 渠道库存查询接口
func TmallSupplychainChannelProductQuantityGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductQuantityGetAPIRequest, resp *fenxiao.TmallSupplychainChannelProductQuantityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
