package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaobanamadpcitemrender 新发商品发布页
// taobao.banamadpc.item.render
//
// 巴拿马供应商通过此接口新发商品发布页
func Taobaobanamadpcitemrender(clt *core.SDKClient, req *product.TaobaobanamadpcitemrenderAPIRequest, session string) (*product.TaobaobanamadpcitemrenderAPIResponse, error) {
	var resp product.TaobaobanamadpcitemrenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
