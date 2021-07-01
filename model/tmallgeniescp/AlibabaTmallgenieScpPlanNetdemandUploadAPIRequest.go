package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest
23-Net Demand净需求回传 API请求
alibaba.tmallgenie.scp.plan.netdemand.upload

Net Demand净需求回传 */
type AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// New
