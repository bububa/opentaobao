package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaItemEditSubmit 商品编辑提交schema信息
// alibaba.item.edit.submit
//
// 商品编辑提交schema信息
func AlibabaItemEditSubmit(clt *core.SDKClient, req *product.AlibabaItemEditSubmitAPIRequest, session string) (*product.AlibabaItemEditSubmitAPIResponse, error) {
	var resp product.AlibabaItemEditSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
