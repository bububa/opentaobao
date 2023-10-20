package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseDeleteproduct 批量删除橱窗商品
// alibaba.scbp.showcase.deleteproduct
//
// 批量删除橱窗商品
func AlibabaScbpShowcaseDeleteproduct(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseDeleteproductAPIRequest, resp *icbushowcase.AlibabaScbpShowcaseDeleteproductAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
