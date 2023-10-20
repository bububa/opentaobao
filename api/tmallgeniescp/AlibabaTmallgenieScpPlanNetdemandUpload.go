package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanNetdemandUpload 23-Net Demand净需求回传
// alibaba.tmallgenie.scp.plan.netdemand.upload
//
// Net Demand净需求回传
func AlibabaTmallgenieScpPlanNetdemandUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
