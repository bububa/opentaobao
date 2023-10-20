package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaobanamadpcitemadd 新发商品
// taobao.banamadpc.item.add
//
// 巴拿马供应商通过此接口新发商品
func Taobaobanamadpcitemadd(clt *core.SDKClient, req *product.TaobaobanamadpcitemaddAPIRequest, session string) (*product.TaobaobanamadpcitemaddAPIResponse, error) {
	var resp product.TaobaobanamadpcitemaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
