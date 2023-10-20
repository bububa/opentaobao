package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemstoreschemaupdate 天猫门店商品编辑
// tmall.item.store.schema.update
//
// 天猫门店商品编辑
func Tmallitemstoreschemaupdate(clt *core.SDKClient, req *product.TmallitemstoreschemaupdateAPIRequest, session string) (*product.TmallitemstoreschemaupdateAPIResponse, error) {
	var resp product.TmallitemstoreschemaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
