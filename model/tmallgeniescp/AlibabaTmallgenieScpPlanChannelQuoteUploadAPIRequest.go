package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest
9.1-同步渠道配额 API请求
alibaba.tmallgenie.scp.plan.channel.quote.upload

同步渠道配额 */
type AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// New
