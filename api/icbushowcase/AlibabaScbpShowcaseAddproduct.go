package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseAddproduct 批量添加橱窗商品
// alibaba.scbp.showcase.addproduct
//
// 批量添加商品到橱窗
func AlibabaScbpShowcaseAddproduct(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseAddproductAPIRequest, resp *icbushowcase.AlibabaScbpShowcaseAddproductAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
