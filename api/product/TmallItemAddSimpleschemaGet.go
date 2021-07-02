package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemAddSimpleschemaGet 天猫发布商品规则获取
// tmall.item.add.simpleschema.get
//
// 通过商家信息获取商品发布字段和规则。
func TmallItemAddSimpleschemaGet(clt *core.SDKClient, req *product.TmallItemAddSimpleschemaGetAPIRequest, session string) (*product.TmallItemAddSimpleschemaGetAPIResponse, error) {
	var resp product.TmallItemAddSimpleschemaGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
