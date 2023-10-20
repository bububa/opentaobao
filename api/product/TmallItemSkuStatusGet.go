package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemskustatusget 商品sku上下架查询
// tmall.item.sku.status.get
//
// 商品sku上下架状态查询
func Tmallitemskustatusget(clt *core.SDKClient, req *product.TmallitemskustatusgetAPIRequest, session string) (*product.TmallitemskustatusgetAPIResponse, error) {
	var resp product.TmallitemskustatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
