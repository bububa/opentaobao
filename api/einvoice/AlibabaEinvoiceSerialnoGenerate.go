package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceSerialnoGenerate 获取统一开票流水号
// alibaba.einvoice.serialno.generate
//
// erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
func AlibabaEinvoiceSerialnoGenerate(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceSerialnoGenerateAPIRequest, resp *einvoice.AlibabaEinvoiceSerialnoGenerateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
