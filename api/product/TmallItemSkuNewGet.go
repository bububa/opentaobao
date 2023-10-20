package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuNewGet 查询sku销售属性标新信息
// tmall.item.sku.new.get
//
// 查询sku销售属性标新信息
func TmallItemSkuNewGet(clt *core.SDKClient, req *product.TmallItemSkuNewGetAPIRequest, session string) (*product.TmallItemSkuNewGetAPIResponse, error) {
	var resp product.TmallItemSkuNewGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
