package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAppRiskScanAPIRequest
应用风险扫描提交接口 API请求
alibaba.security.jaq.app.risk.scan

提交应用进行风险扫描(含漏洞扫描、恶意代码检测、仿冒监测),扫描完成后可通过对应的查询接口查询扫描结果 */
type AlibabaSecurityJaqAppRiskScanAPIRequest struct {
	model.Params
	// 应用信息
	_appInfo *ScanAppInfo
	// 扫描类型：vuln-漏洞扫描 malware-恶意代码检测 fake-仿冒监测 plugin-插件扫描 注: dataType为2时 不支持 仿冒监测
	_scanTypes []string
	// 额外的信息，根据具体业务定
	_extParam string
}

// New
