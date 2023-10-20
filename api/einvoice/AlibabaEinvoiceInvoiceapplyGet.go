package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceinvoiceapplyget 获取商家的开票申请
// alibaba.einvoice.invoiceapply.get
//
// 开票服务商接收到商家发起的开票申请消息后，调用此接口拉取商家详细的开票申请内容
func Alibabaeinvoiceinvoiceapplyget(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceinvoiceapplygetAPIRequest, session string) (*einvoice.AlibabaeinvoiceinvoiceapplygetAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceinvoiceapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
