package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAppShieldAPIRequest
应用加固接口 API请求
alibaba.security.jaq.app.shield

提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果 */
type AlibabaSecurityJaqAppShieldAPIRequest struct {
	model.Params
	// 待加固的应用信息
	_appInfo *ScanAppInfo
	// 渠道列表,多渠道加固时填写
	_channel *ShieldChannel
}

// New
