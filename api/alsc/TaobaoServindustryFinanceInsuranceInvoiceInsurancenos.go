package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoServindustryFinanceInsuranceInvoiceInsurancenos 商家查询本次开票的保险单号
// taobao.servindustry.finance.insurance.invoice.insurancenos
//
// 商家查询本次开票的保险单号
func TaobaoServindustryFinanceInsuranceInvoiceInsurancenos(clt *core.SDKClient, req *alsc.TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest, session string) (*alsc.TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse, error) {
	var resp alsc.TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
