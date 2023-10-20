package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanLocationQuoteUpload 9.2-同步地点配额
// alibaba.tmallgenie.scp.plan.location.quote.upload
//
// 同步地点配额
func AlibabaTmallgenieScpPlanLocationQuoteUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
