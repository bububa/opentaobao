package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemupdatesimpleschemaget 官网同购编辑商品的get接口
// tmall.item.update.simpleschema.get
//
// 官网同购编辑商品的get接口
func Tmallitemupdatesimpleschemaget(clt *core.SDKClient, req *product.TmallitemupdatesimpleschemagetAPIRequest, session string) (*product.TmallitemupdatesimpleschemagetAPIResponse, error) {
	var resp product.TmallitemupdatesimpleschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
