package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest
应用风险详细信息批量查询接口 API请求
alibaba.security.jaq.app.riskdetailbatch.get

用户通过alibaba.security.jaq.app.risk.scanbatch接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息 */
type AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
	// 本地化语言信息,用于指定返回结果内容所使用的语言(默认为zh_CN,目前仅支持zh_CN)
	_locale *Locale
}

// New
