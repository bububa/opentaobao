package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaDataItemGet 获取商品
// alibaba.data.item.get
//
// 获取商品信息，作为客户端Weex鉴权的虚拟api
func AlibabaDataItemGet(clt *core.SDKClient, req *shop.AlibabaDataItemGetAPIRequest, resp *shop.AlibabaDataItemGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
