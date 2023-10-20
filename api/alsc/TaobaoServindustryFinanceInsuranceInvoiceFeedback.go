package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Taobaoservindustryfinanceinsuranceinvoicefeedback 保险-开票结果反馈
// taobao.servindustry.finance.insurance.invoice.feedback
//
// 保险-开票结果反馈
func Taobaoservindustryfinanceinsuranceinvoicefeedback(clt *core.SDKClient, req *alsc.TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest, session string) (*alsc.TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIResponse, error) {
	var resp alsc.TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
