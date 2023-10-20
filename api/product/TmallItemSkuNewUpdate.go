package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemskunewupdate 更新sku销售属性标新状态
// tmall.item.sku.new.update
//
// 更新sku销售属性标新状态
func Tmallitemskunewupdate(clt *core.SDKClient, req *product.TmallitemskunewupdateAPIRequest, session string) (*product.TmallitemskunewupdateAPIResponse, error) {
	var resp product.TmallitemskunewupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
