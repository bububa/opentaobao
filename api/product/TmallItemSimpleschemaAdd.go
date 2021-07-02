package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSimpleschemaAdd 天猫简化发布商品
// tmall.item.simpleschema.add
//
// 天猫简化版schema发布商品。
func TmallItemSimpleschemaAdd(clt *core.SDKClient, req *product.TmallItemSimpleschemaAddAPIRequest, session string) (*product.TmallItemSimpleschemaAddAPIResponse, error) {
	var resp product.TmallItemSimpleschemaAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
