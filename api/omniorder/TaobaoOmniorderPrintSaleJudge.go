package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderPrintSaleJudge 导购员判断
// taobao.omniorder.print.sale.judge
//
// 用于判断当前子账号是否导购员
func TaobaoOmniorderPrintSaleJudge(clt *core.SDKClient, req *omniorder.TaobaoOmniorderPrintSaleJudgeAPIRequest, resp *omniorder.TaobaoOmniorderPrintSaleJudgeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
