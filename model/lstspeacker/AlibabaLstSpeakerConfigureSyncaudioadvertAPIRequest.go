package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest 同步广告 API请求
// alibaba.lst.speaker.configure.syncaudioadvert
//
// 如意音箱广告同步
type AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 音频参数
	_speakerConfigParam4SyncAudioAdvert *SpeakerConfigParam4SyncAudioAdvert
}

// NewAlibabaLstSpeakerConfigureSyncaudioadvertRequest 初始化AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest对象
func NewAlibabaLstSpeakerConfigureSyncaudioadvertRequest() *AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest {
	return &AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.configure.syncaudioadvert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// Get DeviceCode Getter
func (r AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// Set is SpeakerConfigParam4SyncAudioAdvert Setter
// 音频参数
func (r *AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest) SetSpeakerConfigParam4SyncAudioAdvert(_speakerConfigParam4SyncAudioAdvert *SpeakerConfigParam4SyncAudioAdvert) error {
	r._speakerConfigParam4SyncAudioAdvert = _speakerConfigParam4SyncAudioAdvert
	r.Set("speaker_config_param4_sync_audio_advert", _speakerConfigParam4SyncAudioAdvert)
	return nil
}

// Get SpeakerConfigParam4SyncAudioAdvert Getter
func (r AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest) GetSpeakerConfigParam4SyncAudioAdvert() *SpeakerConfigParam4SyncAudioAdvert {
	return r._speakerConfigParam4SyncAudioAdvert
}
