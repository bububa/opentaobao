package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest
变更认证回调 API请求
alibaba.nazca.auth.changeauthapply.callback

变更认证回调 */
type AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest struct {
	model.Params
	// 变更认证回调参数
	_paramChangeAuthCallBackDo *ChangeAuthCallBackDo
}

// New
