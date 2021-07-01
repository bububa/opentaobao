package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthRefreshAPIRequest
刷新token API请求
alibaba.ailabs.tmallgenie.auth.refresh

通过此接口刷新天猫精灵授权token */
type AlibabaAilabsTmallgenieAuthRefreshAPIRequest struct {
	model.Params
	// refresh_token_request
	_refreshTokenRequest *TopRefreshReqDto
}

// New
