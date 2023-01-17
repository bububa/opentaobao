package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqResourceFetchAPIRequest 获取资源文件 API请求
// alibaba.security.jaq.resource.fetch
//
// 在前向化验证流程中提供资源文件服务
type AlibabaSecurityJaqResourceFetchAPIRequest struct {
	model.Params
	// 设备类型可能值有：android ios wp
	_deviceType string
	// 分辨率
	_dpi string
	// 语言类型 zh_CN en_US
	_lang string
}

// NewAlibabaSecurityJaqResourceFetchRequest 初始化AlibabaSecurityJaqResourceFetchAPIRequest对象
func NewAlibabaSecurityJaqResourceFetchRequest() *AlibabaSecurityJaqResourceFetchAPIRequest {
	return &AlibabaSecurityJaqResourceFetchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.resource.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceType is DeviceType Setter
// 设备类型可能值有：android ios wp
func (r *AlibabaSecurityJaqResourceFetchAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetDpi is Dpi Setter
// 分辨率
func (r *AlibabaSecurityJaqResourceFetchAPIRequest) SetDpi(_dpi string) error {
	r._dpi = _dpi
	r.Set("dpi", _dpi)
	return nil
}

// GetDpi Dpi Getter
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetDpi() string {
	return r._dpi
}

// SetLang is Lang Setter
// 语言类型 zh_CN en_US
func (r *AlibabaSecurityJaqResourceFetchAPIRequest) SetLang(_lang string) error {
	r._lang = _lang
	r.Set("lang", _lang)
	return nil
}

// GetLang Lang Getter
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetLang() string {
	return r._lang
}
