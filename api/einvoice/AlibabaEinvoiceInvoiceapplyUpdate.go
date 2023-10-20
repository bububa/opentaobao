package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceinvoiceapplyupdate 商家开票申请单状态回传
// alibaba.einvoice.invoiceapply.update
//
// 开票服务商更新商家开票申请单状态
func Alibabaeinvoiceinvoiceapplyupdate(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceinvoiceapplyupdateAPIRequest, session string) (*einvoice.AlibabaeinvoiceinvoiceapplyupdateAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceinvoiceapplyupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
