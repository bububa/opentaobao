package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpCloudRphit 实人认证云服务日志打点
// alibaba.security.jaq.rp.cloud.rphit
//
// 聚安全实人认证日志打点接口
func AlibabaSecurityJaqRpCloudRphit(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudRphitAPIRequest, session string) (*security.AlibabaSecurityJaqRpCloudRphitAPIResponse, error) {
	var resp security.AlibabaSecurityJaqRpCloudRphitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
