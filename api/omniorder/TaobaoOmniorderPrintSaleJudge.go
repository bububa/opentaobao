package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderPrintSaleJudge 导购员判断
// taobao.omniorder.print.sale.judge
//
// 用于判断当前子账号是否导购员
func TaobaoOmniorderPrintSaleJudge(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderPrintSaleJudgeAPIRequest, resp *omniorder.TaobaoOmniorderPrintSaleJudgeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
