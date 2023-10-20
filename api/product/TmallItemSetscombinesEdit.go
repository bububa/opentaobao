package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSetscombinesEdit 普通商品转套装商品&套装商品编辑接口
// tmall.item.setscombines.edit
//
// 普通商品转套装商品&amp;套装商品编辑接口
func TmallItemSetscombinesEdit(clt *core.SDKClient, req *product.TmallItemSetscombinesEditAPIRequest, resp *product.TmallItemSetscombinesEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
