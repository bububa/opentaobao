package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemstoreupdateschemaget 天猫门店商品修改规则获取
// tmall.item.store.update.schema.get
//
// 天猫门店商品修改规则获取
func Tmallitemstoreupdateschemaget(clt *core.SDKClient, req *product.TmallitemstoreupdateschemagetAPIRequest, session string) (*product.TmallitemstoreupdateschemagetAPIResponse, error) {
	var resp product.TmallitemstoreupdateschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
