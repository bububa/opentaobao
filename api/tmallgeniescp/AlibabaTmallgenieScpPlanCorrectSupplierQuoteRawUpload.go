package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUpload 同步供应商校准后的配额-二级物料
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload
//
// 同步供应商校准后的配额-二级物料
func AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
