package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceClosereq 关闭开票失败请求（失败列表可重试）
// alibaba.einvoice.closereq
//
// 关闭失败开票请求，避免造成重复开票
func AlibabaEinvoiceClosereq(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceClosereqAPIRequest, resp *einvoice.AlibabaEinvoiceClosereqAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
