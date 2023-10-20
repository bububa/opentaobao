package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductUpshelf 产品上架
// tmall.supplychain.channel.product.upshelf
//
// 上架渠道产品
func TmallSupplychainChannelProductUpshelf(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductUpshelfAPIRequest, resp *fenxiao.TmallSupplychainChannelProductUpshelfAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
