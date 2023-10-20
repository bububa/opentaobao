package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqappriskscanAPIRequest 应用风险扫描提交接口 API请求
// alibaba.security.jaq.app.risk.scan
//
// 提交应用进行风险扫描(含漏洞扫描、恶意代码检测、仿冒监测),扫描完成后可通过对应的查询接口查询扫描结果
type AlibabasecurityjaqappriskscanAPIRequest struct {
	model.Params
	// 扫描类型：vuln-漏洞扫描 malware-恶意代码检测 fake-仿冒监测 plugin-插件扫描 注: dataType为2时 不支持 仿冒监测
	_scanTypes []string
	// 额外的信息，根据具体业务定
	_extParam string
	// 应用信息
	_appInfo *ScanAppInfo
}

// NewAlibabasecurityjaqappriskscanRequest 初始化AlibabasecurityjaqappriskscanAPIRequest对象
func NewAlibabasecurityjaqappriskscanRequest() *AlibabasecurityjaqappriskscanAPIRequest {
	return &AlibabasecurityjaqappriskscanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqappriskscanAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.risk.scan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqappriskscanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqappriskscanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScanTypes is ScanTypes Setter
// 扫描类型：vuln-漏洞扫描 malware-恶意代码检测 fake-仿冒监测 plugin-插件扫描 注: dataType为2时 不支持 仿冒监测
func (r *AlibabasecurityjaqappriskscanAPIRequest) SetScanTypes(_scanTypes []string) error {
	r._scanTypes = _scanTypes
	r.Set("scan_types", _scanTypes)
	return nil
}

// GetScanTypes ScanTypes Getter
func (r AlibabasecurityjaqappriskscanAPIRequest) GetScanTypes() []string {
	return r._scanTypes
}

// SetExtParam is ExtParam Setter
// 额外的信息，根据具体业务定
func (r *AlibabasecurityjaqappriskscanAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// GetExtParam ExtParam Getter
func (r AlibabasecurityjaqappriskscanAPIRequest) GetExtParam() string {
	return r._extParam
}

// SetAppInfo is AppInfo Setter
// 应用信息
func (r *AlibabasecurityjaqappriskscanAPIRequest) SetAppInfo(_appInfo *ScanAppInfo) error {
	r._appInfo = _appInfo
	r.Set("app_info", _appInfo)
	return nil
}

// GetAppInfo AppInfo Getter
func (r AlibabasecurityjaqappriskscanAPIRequest) GetAppInfo() *ScanAppInfo {
	return r._appInfo
}
