package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemsetscombinesedit 普通商品转套装商品&套装商品编辑接口
// tmall.item.setscombines.edit
//
// 普通商品转套装商品&amp;套装商品编辑接口
func Tmallitemsetscombinesedit(clt *core.SDKClient, req *product.TmallitemsetscombineseditAPIRequest, session string) (*product.TmallitemsetscombineseditAPIResponse, error) {
	var resp product.TmallitemsetscombineseditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
