package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanCorrectSupplierQuoteUpload 供应商校准后的配额同步
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload
//
// 供应商校准后的配额同步
func AlibabaTmallgenieScpPlanCorrectSupplierQuoteUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
