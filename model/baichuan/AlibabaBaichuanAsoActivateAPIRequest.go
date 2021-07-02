package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanAsoActivateAPIRequest 设备安装活动激活 API请求
// alibaba.baichuan.aso.activate
//
// 设备安装活动激活
type AlibabaBaichuanAsoActivateAPIRequest struct {
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

// NewAlibabaBaichuanAsoActivateRequest 初始化AlibabaBaichuanAsoActivateAPIRequest对象
func NewAlibabaBaichuanAsoActivateRequest() *AlibabaBaichuanAsoActivateAPIRequest {
	return &AlibabaBaichuanAsoActivateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanAsoActivateAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.aso.activate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanAsoActivateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Source Setter
// 来源
func (r *AlibabaBaichuanAsoActivateAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// Get Source Getter
func (r AlibabaBaichuanAsoActivateAPIRequest) GetSource() string {
	return r._source
}

// Set is AppId Setter
// 1-tmail,2-taobao
func (r *AlibabaBaichuanAsoActivateAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// Get AppId Getter
func (r AlibabaBaichuanAsoActivateAPIRequest) GetAppId() string {
	return r._appId
}

// Set is AppOs Setter
// 1-android,2-ios
func (r *AlibabaBaichuanAsoActivateAPIRequest) SetAppOs(_appOs int64) error {
	r._appOs = _appOs
	r.Set("app_os", _appOs)
	return nil
}

// Get AppOs Getter
func (r AlibabaBaichuanAsoActivateAPIRequest) GetAppOs() int64 {
	return r._appOs
}

// Set is DeviceInfo Setter
// 设备信息，ios为idfa ，android 为imei + imsi
func (r *AlibabaBaichuanAsoActivateAPIRequest) SetDeviceInfo(_deviceInfo *AsoDeviceInfoDo) error {
	r._deviceInfo = _deviceInfo
	r.Set("device_info", _deviceInfo)
	return nil
}

// Get DeviceInfo Getter
func (r AlibabaBaichuanAsoActivateAPIRequest) GetDeviceInfo() *AsoDeviceInfoDo {
	return r._deviceInfo
}
