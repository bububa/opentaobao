package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanCorrectSupplierQuoteUpload 供应商校准后的配额同步
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload
//
// 供应商校准后的配额同步
func AlibabaTmallgenieScpPlanCorrectSupplierQuoteUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
