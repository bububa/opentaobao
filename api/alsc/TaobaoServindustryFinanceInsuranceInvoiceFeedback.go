package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoServindustryFinanceInsuranceInvoiceFeedback 保险-开票结果反馈
// taobao.servindustry.finance.insurance.invoice.feedback
//
// 保险-开票结果反馈
func TaobaoServindustryFinanceInsuranceInvoiceFeedback(ctx context.Context, clt *core.SDKClient, req *alsc.TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest, resp *alsc.TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
