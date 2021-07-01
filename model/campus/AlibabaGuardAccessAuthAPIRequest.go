package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGuardAccessAuthAPIRequest
鉴权 API请求
alibaba.guard.access.auth

刷卡鉴权 */
type AlibabaGuardAccessAuthAPIRequest struct {
	model.Params
	// 请求
	_paramIdentifyAuthDTO *IdentifyAuthDto
}

// New
