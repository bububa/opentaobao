package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMemberIdentityRescindfinishAPIRequest
取消确认 API请求
alibaba.member.identity.rescindfinish

取消确认 */
type AlibabaMemberIdentityRescindfinishAPIRequest struct {
	model.Params
	// 取消确认信息
	_rescindFinish *RescindIdentityFinishRequest
}

// New
