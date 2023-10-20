package baichuan

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaBaichuanAsoActivateAPIRequest) Reset() {
	r._source = ""
	r._appId = ""
	r._appOs = 0
	r._deviceInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanAsoActivateAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.aso.activate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanAsoActivateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBaichuanAsoActivateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSource is Source Setter
// 来源
func (r *AlibabaBaichuanAsoActivateAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaBaichuanAsoActivateAPIRequest) GetSource() string {
	return r._source
}

// SetAppId is AppId Setter
// 1-tmail,2-taobao
func (r *AlibabaBaichuanAsoActivateAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaBaichuanAsoActivateAPIRequest) GetAppId() string {
	return r._appId
}

// SetAppOs is AppOs Setter
// 1-android,2-ios
func (r *AlibabaBaichuanAsoActivateAPIRequest) SetAppOs(_appOs int64) error {
	r._appOs = _appOs
	r.Set("app_os", _appOs)
	return nil
}

// GetAppOs AppOs Getter
func (r AlibabaBaichuanAsoActivateAPIRequest) GetAppOs() int64 {
	return r._appOs
}

// SetDeviceInfo is DeviceInfo Setter
// 设备信息，ios为idfa ，android 为imei + imsi
func (r *AlibabaBaichuanAsoActivateAPIRequest) SetDeviceInfo(_deviceInfo *AsoDeviceInfoDo) error {
	r._deviceInfo = _deviceInfo
	r.Set("device_info", _deviceInfo)
	return nil
}

// GetDeviceInfo DeviceInfo Getter
func (r AlibabaBaichuanAsoActivateAPIRequest) GetDeviceInfo() *AsoDeviceInfoDo {
	return r._deviceInfo
}

var poolAlibabaBaichuanAsoActivateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaBaichuanAsoActivateRequest()
	},
}

// GetAlibabaBaichuanAsoActivateRequest 从 sync.Pool 获取 AlibabaBaichuanAsoActivateAPIRequest
func GetAlibabaBaichuanAsoActivateAPIRequest() *AlibabaBaichuanAsoActivateAPIRequest {
	return poolAlibabaBaichuanAsoActivateAPIRequest.Get().(*AlibabaBaichuanAsoActivateAPIRequest)
}

// ReleaseAlibabaBaichuanAsoActivateAPIRequest 将 AlibabaBaichuanAsoActivateAPIRequest 放入 sync.Pool
func ReleaseAlibabaBaichuanAsoActivateAPIRequest(v *AlibabaBaichuanAsoActivateAPIRequest) {
	v.Reset()
	poolAlibabaBaichuanAsoActivateAPIRequest.Put(v)
}
