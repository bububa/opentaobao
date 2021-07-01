package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanChannelGetAPIRequest
5-IBP同步渠道接口 API请求
alibaba.tmallgenie.scp.plan.channel.get

IBP同步渠道接口 */
type AlibabaTmallgenieScpPlanChannelGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
