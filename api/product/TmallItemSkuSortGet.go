package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuSortGet sku销售属性顺序获取
// tmall.item.sku.sort.get
//
// sku销售属性顺序获取
func TmallItemSkuSortGet(clt *core.SDKClient, req *product.TmallItemSkuSortGetAPIRequest, resp *product.TmallItemSkuSortGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
