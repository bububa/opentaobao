package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaobanamadpcitemeditrender 编辑商品发布页
// taobao.banamadpc.item.edit.render
//
// 巴拿马供应商通过此接口获取编辑商品发布页
func Taobaobanamadpcitemeditrender(clt *core.SDKClient, req *product.TaobaobanamadpcitemeditrenderAPIRequest, session string) (*product.TaobaobanamadpcitemeditrenderAPIResponse, error) {
	var resp product.TaobaobanamadpcitemeditrenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
