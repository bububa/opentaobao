package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpCloudRealnameCheck 验证姓名和证件号
// alibaba.security.jaq.rp.cloud.realname.check
//
// 验证姓名和证件号
func AlibabaSecurityJaqRpCloudRealnameCheck(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest, resp *security.AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
