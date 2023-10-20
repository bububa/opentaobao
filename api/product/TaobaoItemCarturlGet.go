package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemCarturlGet 加购URL获取
// taobao.item.carturl.get
//
// 获取加购URL，支持添加商品到购物车
func TaobaoItemCarturlGet(clt *core.SDKClient, req *product.TaobaoItemCarturlGetAPIRequest, resp *product.TaobaoItemCarturlGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
