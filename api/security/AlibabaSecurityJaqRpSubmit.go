package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpSubmit 聚安全实人认证提交认证接口
// alibaba.security.jaq.rp.submit
//
// 聚安全实人认证提交认证接口
func AlibabaSecurityJaqRpSubmit(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpSubmitAPIRequest, session string) (*security.AlibabaSecurityJaqRpSubmitAPIResponse, error) {
	var resp security.AlibabaSecurityJaqRpSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
