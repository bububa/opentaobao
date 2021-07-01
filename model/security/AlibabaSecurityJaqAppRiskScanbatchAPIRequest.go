package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAppRiskScanbatchAPIRequest
应用风险扫描批量提交接口 API请求
alibaba.security.jaq.app.risk.scanbatch

批量提交应用进行风险扫描(含漏洞扫描、恶意代码检测),扫描完成后可通过对应的查询接口查询扫描结果 */
type AlibabaSecurityJaqAppRiskScanbatchAPIRequest struct {
	model.Params
	// APP信息
	_appInfo *AppInfoBatch
	// 扫描类型
	_scanTypes []string
}

// New
