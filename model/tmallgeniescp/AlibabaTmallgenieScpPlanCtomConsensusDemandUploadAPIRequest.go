package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest
22-C2M共识需求回传接口 API请求
alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload

C2M 共识需求回传接口 */
type AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest struct {
	model.Params
	// 对象
	_c2MConsensusDemandRequest *C2MConsensusDemandRequest
}

// New
