package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemUpdateSimpleschemaGet 官网同购编辑商品的get接口
// tmall.item.update.simpleschema.get
//
// 官网同购编辑商品的get接口
func TmallItemUpdateSimpleschemaGet(clt *core.SDKClient, req *product.TmallItemUpdateSimpleschemaGetAPIRequest, resp *product.TmallItemUpdateSimpleschemaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
