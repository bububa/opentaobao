package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanNetdemandRawUpload 二级物料净需求回传(TL+1)
// alibaba.tmallgenie.scp.plan.netdemand.raw.upload
//
// 二级物料净需求回传(TL+1)
func AlibabaTmallgenieScpPlanNetdemandRawUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandRawUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
