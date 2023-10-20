package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemskustatusupdate 商品sku状态更新
// tmall.item.sku.status.update
//
// 商品sku上下架状态更新
func Tmallitemskustatusupdate(clt *core.SDKClient, req *product.TmallitemskustatusupdateAPIRequest, session string) (*product.TmallitemskustatusupdateAPIResponse, error) {
	var resp product.TmallitemskustatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
