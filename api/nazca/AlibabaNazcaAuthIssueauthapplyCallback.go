package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// AlibabaNazcaAuthIssueauthapplyCallback 出证申请回调
// alibaba.nazca.auth.issueauthapply.callback
//
// 出证申请回调
func AlibabaNazcaAuthIssueauthapplyCallback(clt *core.SDKClient, req *nazca.AlibabaNazcaAuthIssueauthapplyCallbackAPIRequest, resp *nazca.AlibabaNazcaAuthIssueauthapplyCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
