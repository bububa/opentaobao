package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemskusortupdate 商品销售属性排序更新
// tmall.item.sku.sort.update
//
// 商品销售属性排序更新
func Tmallitemskusortupdate(clt *core.SDKClient, req *product.TmallitemskusortupdateAPIRequest, session string) (*product.TmallitemskusortupdateAPIResponse, error) {
	var resp product.TmallitemskusortupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
