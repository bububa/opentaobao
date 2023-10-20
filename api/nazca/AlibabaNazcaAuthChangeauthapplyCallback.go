package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// AlibabaNazcaAuthChangeauthapplyCallback 变更认证回调
// alibaba.nazca.auth.changeauthapply.callback
//
// 变更认证回调
func AlibabaNazcaAuthChangeauthapplyCallback(clt *core.SDKClient, req *nazca.AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest, resp *nazca.AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
