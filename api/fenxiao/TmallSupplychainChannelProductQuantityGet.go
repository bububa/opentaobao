package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductQuantityGet 渠道库存查询接口
// tmall.supplychain.channel.product.quantity.get
//
// 渠道库存查询接口
func TmallSupplychainChannelProductQuantityGet(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductQuantityGetAPIRequest, resp *fenxiao.TmallSupplychainChannelProductQuantityGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
