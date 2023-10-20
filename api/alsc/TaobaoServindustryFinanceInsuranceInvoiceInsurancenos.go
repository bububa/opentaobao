package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Taobaoservindustryfinanceinsuranceinvoiceinsurancenos 商家查询本次开票的保险单号
// taobao.servindustry.finance.insurance.invoice.insurancenos
//
// 商家查询本次开票的保险单号
func Taobaoservindustryfinanceinsuranceinvoiceinsurancenos(clt *core.SDKClient, req *alsc.TaobaoservindustryfinanceinsuranceinvoiceinsurancenosAPIRequest, session string) (*alsc.TaobaoservindustryfinanceinsuranceinvoiceinsurancenosAPIResponse, error) {
	var resp alsc.TaobaoservindustryfinanceinsuranceinvoiceinsurancenosAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
