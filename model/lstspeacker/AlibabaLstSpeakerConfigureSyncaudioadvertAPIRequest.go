package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstspeakerconfiguresyncaudioadvertAPIRequest 同步广告 API请求
// alibaba.lst.speaker.configure.syncaudioadvert
//
// 如意音箱广告同步
type AlibabalstspeakerconfiguresyncaudioadvertAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 音频参数
	_speakerConfigParam4SyncAudioAdvert *SpeakerConfigParam4syncAudioAdvert
}

// NewAlibabalstspeakerconfiguresyncaudioadvertRequest 初始化AlibabalstspeakerconfiguresyncaudioadvertAPIRequest对象
func NewAlibabalstspeakerconfiguresyncaudioadvertRequest() *AlibabalstspeakerconfiguresyncaudioadvertAPIRequest {
	return &AlibabalstspeakerconfiguresyncaudioadvertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstspeakerconfiguresyncaudioadvertAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.configure.syncaudioadvert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstspeakerconfiguresyncaudioadvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstspeakerconfiguresyncaudioadvertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCode is DeviceCode Setter
// 设备编码
func (r *AlibabalstspeakerconfiguresyncaudioadvertAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r AlibabalstspeakerconfiguresyncaudioadvertAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetSpeakerConfigParam4SyncAudioAdvert is SpeakerConfigParam4SyncAudioAdvert Setter
// 音频参数
func (r *AlibabalstspeakerconfiguresyncaudioadvertAPIRequest) SetSpeakerConfigParam4SyncAudioAdvert(_speakerConfigParam4SyncAudioAdvert *SpeakerConfigParam4syncAudioAdvert) error {
	r._speakerConfigParam4SyncAudioAdvert = _speakerConfigParam4SyncAudioAdvert
	r.Set("speaker_config_param4_sync_audio_advert", _speakerConfigParam4SyncAudioAdvert)
	return nil
}

// GetSpeakerConfigParam4SyncAudioAdvert SpeakerConfigParam4SyncAudioAdvert Getter
func (r AlibabalstspeakerconfiguresyncaudioadvertAPIRequest) GetSpeakerConfigParam4SyncAudioAdvert() *SpeakerConfigParam4syncAudioAdvert {
	return r._speakerConfigParam4SyncAudioAdvert
}
