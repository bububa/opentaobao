package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest
9.2-同步地点配额 API请求
alibaba.tmallgenie.scp.plan.location.quote.upload

同步地点配额 */
type AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// New
