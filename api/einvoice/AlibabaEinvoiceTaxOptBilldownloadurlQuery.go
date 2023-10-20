package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicetaxoptbilldownloadurlquery 税筹业务账单文件下载URL查询
// alibaba.einvoice.tax.opt.billdownloadurl.query
//
// 税筹业务账单文件下载的URL查询
func Alibabaeinvoicetaxoptbilldownloadurlquery(clt *core.SDKClient, req *einvoice.AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest, session string) (*einvoice.AlibabaeinvoicetaxoptbilldownloadurlqueryAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicetaxoptbilldownloadurlqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
