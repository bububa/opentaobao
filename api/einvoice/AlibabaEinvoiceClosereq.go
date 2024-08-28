package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceClosereq 关闭开票失败请求（失败列表可重试）
// alibaba.einvoice.closereq
//
// 关闭失败开票请求，避免造成重复开票
func AlibabaEinvoiceClosereq(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceClosereqAPIRequest, resp *einvoice.AlibabaEinvoiceClosereqAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
