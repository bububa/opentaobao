package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemskusortget sku销售属性顺序获取
// tmall.item.sku.sort.get
//
// sku销售属性顺序获取
func Tmallitemskusortget(clt *core.SDKClient, req *product.TmallitemskusortgetAPIRequest, session string) (*product.TmallitemskusortgetAPIResponse, error) {
	var resp product.TmallitemskusortgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
