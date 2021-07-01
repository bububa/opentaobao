package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMemberIdentitySignfinishAPIRequest
签约确认 API请求
alibaba.member.identity.signfinish

签约确认 */
type AlibabaMemberIdentitySignfinishAPIRequest struct {
	model.Params
	// 签约确认信息
	_signFinish *SignIdentityFinishRequest
}

// New
