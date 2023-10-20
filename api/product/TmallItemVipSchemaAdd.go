package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemvipschemaadd 大商家商品发布接口
// tmall.item.vip.schema.add
//
// 大商家商品发布接口
func Tmallitemvipschemaadd(clt *core.SDKClient, req *product.TmallitemvipschemaaddAPIRequest, session string) (*product.TmallitemvipschemaaddAPIResponse, error) {
	var resp product.TmallitemvipschemaaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
