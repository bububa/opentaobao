package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

/* AlibabaSecurityJaqRpRphit
聚安全-实人认证日志打点接口
alibaba.security.jaq.rp.rphit

聚安全实人认证日志打点接口 */
func AlibabaSecurityJaqRpRphit(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpRphitAPIRequest, session string) (*security.AlibabaSecurityJaqRpRphitAPIResponse, error) {
	var resp security.AlibabaSecurityJaqRpRphitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
