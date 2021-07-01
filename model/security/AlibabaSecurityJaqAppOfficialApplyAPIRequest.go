package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAppOfficialApplyAPIRequest
聚安全官方应用申请 API请求
alibaba.security.jaq.app.official.apply

官方应用申请接口 */
type AlibabaSecurityJaqAppOfficialApplyAPIRequest struct {
	model.Params
	// 官方应用申请入参
	_officialAppApplyRequest *OfficialAppApplyRequest
}

// New
