package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemvipaddschemaget vip商家发布商品的获取规则接口
// tmall.item.vip.add.schema.get
//
// 获取vip商家发布商品的规则
func Tmallitemvipaddschemaget(clt *core.SDKClient, req *product.TmallitemvipaddschemagetAPIRequest, session string) (*product.TmallitemvipaddschemagetAPIResponse, error) {
	var resp product.TmallitemvipaddschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
