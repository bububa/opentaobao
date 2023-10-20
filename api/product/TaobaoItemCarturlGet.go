package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaoitemcarturlget 加购URL获取
// taobao.item.carturl.get
//
// 获取加购URL，支持添加商品到购物车
func Taobaoitemcarturlget(clt *core.SDKClient, req *product.TaobaoitemcarturlgetAPIRequest, session string) (*product.TaobaoitemcarturlgetAPIResponse, error) {
	var resp product.TaobaoitemcarturlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
