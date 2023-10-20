package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppOfficialApply 聚安全官方应用申请
// alibaba.security.jaq.app.official.apply
//
// 官方应用申请接口
func AlibabaSecurityJaqAppOfficialApply(clt *core.SDKClient, req *security.AlibabaSecurityJaqAppOfficialApplyAPIRequest, session string) (*security.AlibabaSecurityJaqAppOfficialApplyAPIResponse, error) {
	var resp security.AlibabaSecurityJaqAppOfficialApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
