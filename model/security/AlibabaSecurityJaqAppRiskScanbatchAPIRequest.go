package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppRiskScanbatchAPIRequest 应用风险扫描批量提交接口 API请求
// alibaba.security.jaq.app.risk.scanbatch
//
// 批量提交应用进行风险扫描(含漏洞扫描、恶意代码检测),扫描完成后可通过对应的查询接口查询扫描结果
type AlibabaSecurityJaqAppRiskScanbatchAPIRequest struct {
	model.Params
	// 扫描类型
	_scanTypes []string
	// APP信息
	_appInfo *AppInfoBatch
}

// NewAlibabaSecurityJaqAppRiskScanbatchRequest 初始化AlibabaSecurityJaqAppRiskScanbatchAPIRequest对象
func NewAlibabaSecurityJaqAppRiskScanbatchRequest() *AlibabaSecurityJaqAppRiskScanbatchAPIRequest {
	return &AlibabaSecurityJaqAppRiskScanbatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppRiskScanbatchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.risk.scanbatch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppRiskScanbatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqAppRiskScanbatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScanTypes is ScanTypes Setter
// 扫描类型
func (r *AlibabaSecurityJaqAppRiskScanbatchAPIRequest) SetScanTypes(_scanTypes []string) error {
	r._scanTypes = _scanTypes
	r.Set("scan_types", _scanTypes)
	return nil
}

// GetScanTypes ScanTypes Getter
func (r AlibabaSecurityJaqAppRiskScanbatchAPIRequest) GetScanTypes() []string {
	return r._scanTypes
}

// SetAppInfo is AppInfo Setter
// APP信息
func (r *AlibabaSecurityJaqAppRiskScanbatchAPIRequest) SetAppInfo(_appInfo *AppInfoBatch) error {
	r._appInfo = _appInfo
	r.Set("app_info", _appInfo)
	return nil
}

// GetAppInfo AppInfo Getter
func (r AlibabaSecurityJaqAppRiskScanbatchAPIRequest) GetAppInfo() *AppInfoBatch {
	return r._appInfo
}
