package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuNewUpdate 更新sku销售属性标新状态
// tmall.item.sku.new.update
//
// 更新sku销售属性标新状态
func TmallItemSkuNewUpdate(clt *core.SDKClient, req *product.TmallItemSkuNewUpdateAPIRequest, session string) (*product.TmallItemSkuNewUpdateAPIResponse, error) {
	var resp product.TmallItemSkuNewUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
