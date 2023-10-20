package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicetaxoptsalaryrequestsingleaccept 单明细发薪受理
// alibaba.einvoice.tax.opt.salaryrequest.singleaccept
//
// 单明细发薪受理
func Alibabaeinvoicetaxoptsalaryrequestsingleaccept(clt *core.SDKClient, req *einvoice.AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest, session string) (*einvoice.AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
