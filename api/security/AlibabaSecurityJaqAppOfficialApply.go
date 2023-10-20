package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppOfficialApply 聚安全官方应用申请
// alibaba.security.jaq.app.official.apply
//
// 官方应用申请接口
func AlibabaSecurityJaqAppOfficialApply(clt *core.SDKClient, req *security.AlibabaSecurityJaqAppOfficialApplyAPIRequest, resp *security.AlibabaSecurityJaqAppOfficialApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
