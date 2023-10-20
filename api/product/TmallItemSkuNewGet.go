package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemskunewget 查询sku销售属性标新信息
// tmall.item.sku.new.get
//
// 查询sku销售属性标新信息
func Tmallitemskunewget(clt *core.SDKClient, req *product.TmallitemskunewgetAPIRequest, session string) (*product.TmallitemskunewgetAPIResponse, error) {
	var resp product.TmallitemskunewgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
