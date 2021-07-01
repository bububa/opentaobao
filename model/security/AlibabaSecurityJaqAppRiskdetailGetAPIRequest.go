package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAppRiskdetailGetAPIRequest
应用风险详细信息查询接口 API请求
alibaba.security.jaq.app.riskdetail.get

用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息 */
type AlibabaSecurityJaqAppRiskdetailGetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
	// 本地化语言信息
	_locale *Locale
}

// New
