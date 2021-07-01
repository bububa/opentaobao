package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

/* AlibabaSecurityJaqRpCloudStart
实人认证云开始认证
alibaba.security.jaq.rp.cloud.start

聚安全实人认证开始 */
func AlibabaSecurityJaqRpCloudStart(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudStartAPIRequest, session string) (*security.AlibabaSecurityJaqRpCloudStartAPIResponse, error) {
	var resp security.AlibabaSecurityJaqRpCloudStartAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
