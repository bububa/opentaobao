package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceApplyGet 开票申请数据获取接口
// alibaba.einvoice.apply.get
//
// ERP获取开票申请数据
func AlibabaEinvoiceApplyGet(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceApplyGetAPIRequest, resp *einvoice.AlibabaEinvoiceApplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
