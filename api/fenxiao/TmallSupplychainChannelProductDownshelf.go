package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductDownshelf 产品下架
// tmall.supplychain.channel.product.downshelf
//
// 产品下架
func TmallSupplychainChannelProductDownshelf(ctx context.Context, clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductDownshelfAPIRequest, resp *fenxiao.TmallSupplychainChannelProductDownshelfAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
