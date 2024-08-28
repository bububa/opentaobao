package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoicePaperCommonReturn 纸票通用回传接口
// alibaba.einvoice.paper.common.return
//
// 纸票通用回传接口（打印回传、注册回传等），只返回成功or失败
func AlibabaEinvoicePaperCommonReturn(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoicePaperCommonReturnAPIRequest, resp *einvoice.AlibabaEinvoicePaperCommonReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
