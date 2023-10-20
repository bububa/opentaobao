package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUpload 同步供应商校准后的配额-二级物料
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload
//
// 同步供应商校准后的配额-二级物料
func AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
