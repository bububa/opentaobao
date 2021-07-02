package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseAddproduct 批量添加橱窗商品
// alibaba.scbp.showcase.addproduct
//
// 批量添加商品到橱窗
func AlibabaScbpShowcaseAddproduct(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseAddproductAPIRequest, session string) (*icbushowcase.AlibabaScbpShowcaseAddproductAPIResponse, error) {
	var resp icbushowcase.AlibabaScbpShowcaseAddproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
