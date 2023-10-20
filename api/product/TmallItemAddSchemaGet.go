package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemaddschemaget 天猫发布商品规则获取
// tmall.item.add.schema.get
//
// 通过类目以及productId获取商品发布规则；
func Tmallitemaddschemaget(clt *core.SDKClient, req *product.TmallitemaddschemagetAPIRequest, session string) (*product.TmallitemaddschemagetAPIResponse, error) {
	var resp product.TmallitemaddschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
