package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanSkuSupplierQuoteUpload 标准供应商配额同步
// alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload
//
// 标准供应商配额同步
func AlibabaTmallgenieScpPlanSkuSupplierQuoteUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
