package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuanasoqueryAPIRequest 查询app在设备上的安装信息 API请求
// alibaba.baichuan.aso.query
//
// 查询app在设备上的安装信息
type AlibababaichuanasoqueryAPIRequest struct {
	model.Params
	// 设备信息，ios为idfa ，android 为imei + imsi
	_deviceInfoList []AsoDeviceInfoDo
	// 1-tmail,2-taobao
	_appId string
	// 1-android,2-ios
	_appOs int64
}

// NewAlibababaichuanasoqueryRequest 初始化AlibababaichuanasoqueryAPIRequest对象
func NewAlibababaichuanasoqueryRequest() *AlibababaichuanasoqueryAPIRequest {
	return &AlibababaichuanasoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababaichuanasoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.aso.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababaichuanasoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababaichuanasoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceInfoList is DeviceInfoList Setter
// 设备信息，ios为idfa ，android 为imei + imsi
func (r *AlibababaichuanasoqueryAPIRequest) SetDeviceInfoList(_deviceInfoList []AsoDeviceInfoDo) error {
	r._deviceInfoList = _deviceInfoList
	r.Set("device_info_list", _deviceInfoList)
	return nil
}

// GetDeviceInfoList DeviceInfoList Getter
func (r AlibababaichuanasoqueryAPIRequest) GetDeviceInfoList() []AsoDeviceInfoDo {
	return r._deviceInfoList
}

// SetAppId is AppId Setter
// 1-tmail,2-taobao
func (r *AlibababaichuanasoqueryAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibababaichuanasoqueryAPIRequest) GetAppId() string {
	return r._appId
}

// SetAppOs is AppOs Setter
// 1-android,2-ios
func (r *AlibababaichuanasoqueryAPIRequest) SetAppOs(_appOs int64) error {
	r._appOs = _appOs
	r.Set("app_os", _appOs)
	return nil
}

// GetAppOs AppOs Getter
func (r AlibababaichuanasoqueryAPIRequest) GetAppOs() int64 {
	return r._appOs
}
