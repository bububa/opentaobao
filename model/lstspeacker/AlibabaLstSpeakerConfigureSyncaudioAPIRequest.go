package lstspeacker

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerConfigureSyncaudioAPIRequest 音频同步 API请求
// alibaba.lst.speaker.configure.syncaudio
//
// 音频同步
type AlibabaLstSpeakerConfigureSyncaudioAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 参数
	_speakerConfigParam4SyncAudio *SpeakerConfigParam4SyncAudio
}

// NewAlibabaLstSpeakerConfigureSyncaudioRequest 初始化AlibabaLstSpeakerConfigureSyncaudioAPIRequest对象
func NewAlibabaLstSpeakerConfigureSyncaudioRequest() *AlibabaLstSpeakerConfigureSyncaudioAPIRequest {
	return &AlibabaLstSpeakerConfigureSyncaudioAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstSpeakerConfigureSyncaudioAPIRequest) Reset() {
	r._deviceCode = ""
	r._speakerConfigParam4SyncAudio = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerConfigureSyncaudioAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.configure.syncaudio"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerConfigureSyncaudioAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstSpeakerConfigureSyncaudioAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCode is DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerConfigureSyncaudioAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r AlibabaLstSpeakerConfigureSyncaudioAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetSpeakerConfigParam4SyncAudio is SpeakerConfigParam4SyncAudio Setter
// 参数
func (r *AlibabaLstSpeakerConfigureSyncaudioAPIRequest) SetSpeakerConfigParam4SyncAudio(_speakerConfigParam4SyncAudio *SpeakerConfigParam4SyncAudio) error {
	r._speakerConfigParam4SyncAudio = _speakerConfigParam4SyncAudio
	r.Set("speaker_config_param4_sync_audio", _speakerConfigParam4SyncAudio)
	return nil
}

// GetSpeakerConfigParam4SyncAudio SpeakerConfigParam4SyncAudio Getter
func (r AlibabaLstSpeakerConfigureSyncaudioAPIRequest) GetSpeakerConfigParam4SyncAudio() *SpeakerConfigParam4SyncAudio {
	return r._speakerConfigParam4SyncAudio
}

var poolAlibabaLstSpeakerConfigureSyncaudioAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstSpeakerConfigureSyncaudioRequest()
	},
}

// GetAlibabaLstSpeakerConfigureSyncaudioRequest 从 sync.Pool 获取 AlibabaLstSpeakerConfigureSyncaudioAPIRequest
func GetAlibabaLstSpeakerConfigureSyncaudioAPIRequest() *AlibabaLstSpeakerConfigureSyncaudioAPIRequest {
	return poolAlibabaLstSpeakerConfigureSyncaudioAPIRequest.Get().(*AlibabaLstSpeakerConfigureSyncaudioAPIRequest)
}

// ReleaseAlibabaLstSpeakerConfigureSyncaudioAPIRequest 将 AlibabaLstSpeakerConfigureSyncaudioAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstSpeakerConfigureSyncaudioAPIRequest(v *AlibabaLstSpeakerConfigureSyncaudioAPIRequest) {
	v.Reset()
	poolAlibabaLstSpeakerConfigureSyncaudioAPIRequest.Put(v)
}
