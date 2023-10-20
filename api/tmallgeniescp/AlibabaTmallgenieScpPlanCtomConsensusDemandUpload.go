package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanCtomConsensusDemandUpload 22-C2M共识需求回传接口
// alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload
//
// C2M 共识需求回传接口
func AlibabaTmallgenieScpPlanCtomConsensusDemandUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
