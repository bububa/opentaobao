package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemvipschemaupdate 大商家商品编辑接口
// tmall.item.vip.schema.update
//
// 大商家编辑商品
func Tmallitemvipschemaupdate(clt *core.SDKClient, req *product.TmallitemvipschemaupdateAPIRequest, session string) (*product.TmallitemvipschemaupdateAPIResponse, error) {
	var resp product.TmallitemvipschemaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
