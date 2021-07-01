package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAppOfficialVerifyAPIRequest
聚安全验证官方应用接口 API请求
alibaba.security.jaq.app.official.verify

接入用户来查询应用是否为官方应用 */
type AlibabaSecurityJaqAppOfficialVerifyAPIRequest struct {
	model.Params
	// 验证参数
	_officialAppVerifyRequest *OfficialAppVerifyRequest
}

// New
