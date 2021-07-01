package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaAuthAuthapplyCallbackAPIRequest
认证的统一回调接口 API请求
alibaba.nazca.auth.authapply.callback

认证的统一回调接口 */
type AlibabaNazcaAuthAuthapplyCallbackAPIRequest struct {
	model.Params
	// 认证回调参数
	_authApplyDoneCallbackDo *AuthApplyDoneCallBackDo
}

// New
