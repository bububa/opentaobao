package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* AlibabaTmallgenieScpPlanNetdemandUpload
23-Net Demand净需求回传
alibaba.tmallgenie.scp.plan.netdemand.upload

Net Demand净需求回传 */
func AlibabaTmallgenieScpPlanNetdemandUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandUploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
