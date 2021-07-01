package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallItemUpdateSimpleschemaGet
官网同购编辑商品的get接口
tmall.item.update.simpleschema.get

官网同购编辑商品的get接口 */
func TmallItemUpdateSimpleschemaGet(clt *core.SDKClient, req *product.TmallItemUpdateSimpleschemaGetAPIRequest, session string) (*product.TmallItemUpdateSimpleschemaGetAPIResponse, error) {
	var resp product.TmallItemUpdateSimpleschemaGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
