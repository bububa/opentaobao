package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqLoginpreventionResultFetch 获取登录保护结果
// alibaba.security.jaq.loginprevention.result.fetch
//
// 获取登录保护结果
func AlibabaSecurityJaqLoginpreventionResultFetch(clt *core.SDKClient, req *security.AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest, resp *security.AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
