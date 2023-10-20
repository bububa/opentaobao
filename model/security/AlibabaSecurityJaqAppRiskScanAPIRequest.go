package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppRiskScanAPIRequest 应用风险扫描提交接口 API请求
// alibaba.security.jaq.app.risk.scan
//
// 提交应用进行风险扫描(含漏洞扫描、恶意代码检测、仿冒监测),扫描完成后可通过对应的查询接口查询扫描结果
type AlibabaSecurityJaqAppRiskScanAPIRequest struct {
	model.Params
	// 扫描类型：vuln-漏洞扫描 malware-恶意代码检测 fake-仿冒监测 plugin-插件扫描 注: dataType为2时 不支持 仿冒监测
	_scanTypes []string
	// 额外的信息，根据具体业务定
	_extParam string
	// 应用信息
	_appInfo *ScanAppInfo
}

// NewAlibabaSecurityJaqAppRiskScanRequest 初始化AlibabaSecurityJaqAppRiskScanAPIRequest对象
func NewAlibabaSecurityJaqAppRiskScanRequest() *AlibabaSecurityJaqAppRiskScanAPIRequest {
	return &AlibabaSecurityJaqAppRiskScanAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqAppRiskScanAPIRequest) Reset() {
	r._scanTypes = r._scanTypes[:0]
	r._extParam = ""
	r._appInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.risk.scan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScanTypes is ScanTypes Setter
// 扫描类型：vuln-漏洞扫描 malware-恶意代码检测 fake-仿冒监测 plugin-插件扫描 注: dataType为2时 不支持 仿冒监测
func (r *AlibabaSecurityJaqAppRiskScanAPIRequest) SetScanTypes(_scanTypes []string) error {
	r._scanTypes = _scanTypes
	r.Set("scan_types", _scanTypes)
	return nil
}

// GetScanTypes ScanTypes Getter
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetScanTypes() []string {
	return r._scanTypes
}

// SetExtParam is ExtParam Setter
// 额外的信息，根据具体业务定
func (r *AlibabaSecurityJaqAppRiskScanAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// GetExtParam ExtParam Getter
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetExtParam() string {
	return r._extParam
}

// SetAppInfo is AppInfo Setter
// 应用信息
func (r *AlibabaSecurityJaqAppRiskScanAPIRequest) SetAppInfo(_appInfo *ScanAppInfo) error {
	r._appInfo = _appInfo
	r.Set("app_info", _appInfo)
	return nil
}

// GetAppInfo AppInfo Getter
func (r AlibabaSecurityJaqAppRiskScanAPIRequest) GetAppInfo() *ScanAppInfo {
	return r._appInfo
}

var poolAlibabaSecurityJaqAppRiskScanAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqAppRiskScanRequest()
	},
}

// GetAlibabaSecurityJaqAppRiskScanRequest 从 sync.Pool 获取 AlibabaSecurityJaqAppRiskScanAPIRequest
func GetAlibabaSecurityJaqAppRiskScanAPIRequest() *AlibabaSecurityJaqAppRiskScanAPIRequest {
	return poolAlibabaSecurityJaqAppRiskScanAPIRequest.Get().(*AlibabaSecurityJaqAppRiskScanAPIRequest)
}

// ReleaseAlibabaSecurityJaqAppRiskScanAPIRequest 将 AlibabaSecurityJaqAppRiskScanAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqAppRiskScanAPIRequest(v *AlibabaSecurityJaqAppRiskScanAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqAppRiskScanAPIRequest.Put(v)
}
