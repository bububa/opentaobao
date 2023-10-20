package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuStatusGet 商品sku上下架查询
// tmall.item.sku.status.get
//
// 商品sku上下架状态查询
func TmallItemSkuStatusGet(clt *core.SDKClient, req *product.TmallItemSkuStatusGetAPIRequest, session string) (*product.TmallItemSkuStatusGetAPIResponse, error) {
	var resp product.TmallItemSkuStatusGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
