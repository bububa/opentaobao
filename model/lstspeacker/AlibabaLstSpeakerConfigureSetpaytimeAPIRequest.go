package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstspeakerconfiguresetpaytimeAPIRequest 音箱播放配置 API请求
// alibaba.lst.speaker.configure.setpaytime
//
// 音箱播放配置
type AlibabalstspeakerconfiguresetpaytimeAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 开始时间
	_playStartTime string
	// 结束时间
	_playEndTime string
	// 是否播放广告
	_isOnlyPlayAdvert bool
	// 是否设置播放时间
	_isSetPlayTime bool
}

// NewAlibabalstspeakerconfiguresetpaytimeRequest 初始化AlibabalstspeakerconfiguresetpaytimeAPIRequest对象
func NewAlibabalstspeakerconfiguresetpaytimeRequest() *AlibabalstspeakerconfiguresetpaytimeAPIRequest {
	return &AlibabalstspeakerconfiguresetpaytimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstspeakerconfiguresetpaytimeAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.configure.setpaytime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstspeakerconfiguresetpaytimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstspeakerconfiguresetpaytimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCode is DeviceCode Setter
// 设备编码
func (r *AlibabalstspeakerconfiguresetpaytimeAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r AlibabalstspeakerconfiguresetpaytimeAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetPlayStartTime is PlayStartTime Setter
// 开始时间
func (r *AlibabalstspeakerconfiguresetpaytimeAPIRequest) SetPlayStartTime(_playStartTime string) error {
	r._playStartTime = _playStartTime
	r.Set("play_start_time", _playStartTime)
	return nil
}

// GetPlayStartTime PlayStartTime Getter
func (r AlibabalstspeakerconfiguresetpaytimeAPIRequest) GetPlayStartTime() string {
	return r._playStartTime
}

// SetPlayEndTime is PlayEndTime Setter
// 结束时间
func (r *AlibabalstspeakerconfiguresetpaytimeAPIRequest) SetPlayEndTime(_playEndTime string) error {
	r._playEndTime = _playEndTime
	r.Set("play_end_time", _playEndTime)
	return nil
}

// GetPlayEndTime PlayEndTime Getter
func (r AlibabalstspeakerconfiguresetpaytimeAPIRequest) GetPlayEndTime() string {
	return r._playEndTime
}

// SetIsOnlyPlayAdvert is IsOnlyPlayAdvert Setter
// 是否播放广告
func (r *AlibabalstspeakerconfiguresetpaytimeAPIRequest) SetIsOnlyPlayAdvert(_isOnlyPlayAdvert bool) error {
	r._isOnlyPlayAdvert = _isOnlyPlayAdvert
	r.Set("is_only_play_advert", _isOnlyPlayAdvert)
	return nil
}

// GetIsOnlyPlayAdvert IsOnlyPlayAdvert Getter
func (r AlibabalstspeakerconfiguresetpaytimeAPIRequest) GetIsOnlyPlayAdvert() bool {
	return r._isOnlyPlayAdvert
}

// SetIsSetPlayTime is IsSetPlayTime Setter
// 是否设置播放时间
func (r *AlibabalstspeakerconfiguresetpaytimeAPIRequest) SetIsSetPlayTime(_isSetPlayTime bool) error {
	r._isSetPlayTime = _isSetPlayTime
	r.Set("is_set_play_time", _isSetPlayTime)
	return nil
}

// GetIsSetPlayTime IsSetPlayTime Getter
func (r AlibabalstspeakerconfiguresetpaytimeAPIRequest) GetIsSetPlayTime() bool {
	return r._isSetPlayTime
}
