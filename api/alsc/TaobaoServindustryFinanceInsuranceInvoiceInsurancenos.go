package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoServindustryFinanceInsuranceInvoiceInsurancenos 商家查询本次开票的保险单号
// taobao.servindustry.finance.insurance.invoice.insurancenos
//
// 商家查询本次开票的保险单号
func TaobaoServindustryFinanceInsuranceInvoiceInsurancenos(ctx context.Context, clt *core.SDKClient, req *alsc.TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest, resp *alsc.TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
