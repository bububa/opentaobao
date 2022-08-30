package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptBilldownloadurlQuery 税筹业务账单文件下载URL查询
// alibaba.einvoice.tax.opt.billdownloadurl.query
//
// 税筹业务账单文件下载的URL查询
func AlibabaEinvoiceTaxOptBilldownloadurlQuery(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIRequest, session string) (*einvoice.AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
