package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoServindustryFinanceInsuranceInvoiceFeedback 保险-开票结果反馈
// taobao.servindustry.finance.insurance.invoice.feedback
//
// 保险-开票结果反馈
func TaobaoServindustryFinanceInsuranceInvoiceFeedback(clt *core.SDKClient, req *alsc.TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest, session string) (*alsc.TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse, error) {
	var resp alsc.TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
