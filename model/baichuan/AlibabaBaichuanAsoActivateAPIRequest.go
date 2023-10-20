package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuanasoactivateAPIRequest 设备安装活动激活 API请求
// alibaba.baichuan.aso.activate
//
// 设备安装活动激活
type AlibababaichuanasoactivateAPIRequest struct {
	model.Params
	// 来源
	_source string
	// 1-tmail,2-taobao
	_appId string
	// 1-android,2-ios
	_appOs int64
	// 设备信息，ios为idfa ，android 为imei + imsi
	_deviceInfo *AsoDeviceInfoDo
}

// NewAlibababaichuanasoactivateRequest 初始化AlibababaichuanasoactivateAPIRequest对象
func NewAlibababaichuanasoactivateRequest() *AlibababaichuanasoactivateAPIRequest {
	return &AlibababaichuanasoactivateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababaichuanasoactivateAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.aso.activate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababaichuanasoactivateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababaichuanasoactivateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSource is Source Setter
// 来源
func (r *AlibababaichuanasoactivateAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibababaichuanasoactivateAPIRequest) GetSource() string {
	return r._source
}

// SetAppId is AppId Setter
// 1-tmail,2-taobao
func (r *AlibababaichuanasoactivateAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibababaichuanasoactivateAPIRequest) GetAppId() string {
	return r._appId
}

// SetAppOs is AppOs Setter
// 1-android,2-ios
func (r *AlibababaichuanasoactivateAPIRequest) SetAppOs(_appOs int64) error {
	r._appOs = _appOs
	r.Set("app_os", _appOs)
	return nil
}

// GetAppOs AppOs Getter
func (r AlibababaichuanasoactivateAPIRequest) GetAppOs() int64 {
	return r._appOs
}

// SetDeviceInfo is DeviceInfo Setter
// 设备信息，ios为idfa ，android 为imei + imsi
func (r *AlibababaichuanasoactivateAPIRequest) SetDeviceInfo(_deviceInfo *AsoDeviceInfoDo) error {
	r._deviceInfo = _deviceInfo
	r.Set("device_info", _deviceInfo)
	return nil
}

// GetDeviceInfo DeviceInfo Getter
func (r AlibababaichuanasoactivateAPIRequest) GetDeviceInfo() *AsoDeviceInfoDo {
	return r._deviceInfo
}
