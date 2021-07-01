package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest
20-IBP共识需求回传接口 API请求
alibaba.tmallgenie.scp.plan.consensus.demand.upload

IBP共识需求回传接口 */
type AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest struct {
	model.Params
	// 入参
	_consensusDemandRequest *ConsensusDemandRequest
}

// New
