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

// NewAlibabaSecurityJaqAppRiskScanRequest 初始化AlibabaSecurityJaqAppRiskScanAPIRequest对象
func NewAlibabaSecurityJaqAppRiskScanRequest() *AlibabaSecurityJaqAppRiskScanAPIRequest {
	return &AlibabaSecurityJaqAppRiskScanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.risk.scan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppInfo Setter
// 应用信息
func (r *AlibabaSecurityJaqAppRiskScanAPIRequest) SetAppInfo(_appInfo *ScanAppInfo) error {
	r._appInfo = _appInfo
	r.Set("app_info", _appInfo)
	return nil
}

// Get AppInfo Getter
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetAppInfo() *ScanAppInfo {
	return r._appInfo
}

// Set is ScanTypes Setter
// 扫描类型：vuln-漏洞扫描 malware-恶意代码检测 fake-仿冒监测 plugin-插件扫描 注: dataType为2时 不支持 仿冒监测
func (r *AlibabaSecurityJaqAppRiskScanAPIRequest) SetScanTypes(_scanTypes []string) error {
	r._scanTypes = _scanTypes
	r.Set("scan_types", _scanTypes)
	return nil
}

// Get ScanTypes Getter
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetScanTypes() []string {
	return r._scanTypes
}

// Set is ExtParam Setter
// 额外的信息，根据具体业务定
func (r *AlibabaSecurityJaqAppRiskScanAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// Get ExtParam Getter
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetExtParam() string {
	return r._extParam
}
