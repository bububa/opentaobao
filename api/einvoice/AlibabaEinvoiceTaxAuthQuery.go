package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxAuthQuery 发票中台授权信息获取
// alibaba.einvoice.tax.auth.query
//
// 发票中台授权信息获取
func AlibabaEinvoiceTaxAuthQuery(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxAuthQueryAPIRequest, session string) (*einvoice.AlibabaEinvoiceTaxAuthQueryAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceTaxAuthQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
