package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemAddSchemaGet 天猫发布商品规则获取
// tmall.item.add.schema.get
//
// 通过类目以及productId获取商品发布规则；
func TmallItemAddSchemaGet(clt *core.SDKClient, req *product.TmallItemAddSchemaGetAPIRequest, session string) (*product.TmallItemAddSchemaGetAPIResponse, error) {
	var resp product.TmallItemAddSchemaGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
