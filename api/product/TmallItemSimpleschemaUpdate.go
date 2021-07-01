package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallItemSimpleschemaUpdate
天猫简化编辑商品
tmall.item.simpleschema.update

国外大商家天猫简化编辑商品 */
func TmallItemSimpleschemaUpdate(clt *core.SDKClient, req *product.TmallItemSimpleschemaUpdateAPIRequest, session string) (*product.TmallItemSimpleschemaUpdateAPIResponse, error) {
	var resp product.TmallItemSimpleschemaUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
