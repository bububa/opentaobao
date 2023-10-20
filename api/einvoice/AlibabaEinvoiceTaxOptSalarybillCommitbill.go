package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicetaxoptsalarybillcommitbill 提交发薪账单
// alibaba.einvoice.tax.opt.salarybill.commitbill
//
// 提交发薪账单
func Alibabaeinvoicetaxoptsalarybillcommitbill(clt *core.SDKClient, req *einvoice.AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest, session string) (*einvoice.AlibabaeinvoicetaxoptsalarybillcommitbillAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicetaxoptsalarybillcommitbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
