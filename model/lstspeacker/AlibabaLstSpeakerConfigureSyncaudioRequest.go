package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
音频同步 API请求
alibaba.lst.speaker.configure.syncaudio

音频同步
*/
type AlibabaLstSpeakerConfigureSyncaudioRequest struct {
    model.Params
    // 设备编码
    _deviceCode   string
    // 参数
    _speakerConfigParam4SyncAudio   *SpeakerConfigParam4SyncAudio
}

// 初始化AlibabaLstSpeakerConfigureSyncaudioRequest对象
func NewAlibabaLstSpeakerConfigureSyncaudioRequest() *AlibabaLstSpeakerConfigureSyncaudioRequest{
    return &AlibabaLstSpeakerConfigureSyncaudioRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerConfigureSyncaudioRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.configure.syncaudio"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerConfigureSyncaudioRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerConfigureSyncaudioRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaLstSpeakerConfigureSyncaudioRequest) GetDeviceCode() string {
    return r._deviceCode
}
// SpeakerConfigParam4SyncAudio Setter
// 参数
func (r *AlibabaLstSpeakerConfigureSyncaudioRequest) SetSpeakerConfigParam4SyncAudio(_speakerConfigParam4SyncAudio *SpeakerConfigParam4SyncAudio) error {
    r._speakerConfigParam4SyncAudio = _speakerConfigParam4SyncAudio
    r.Set("speaker_config_param4_sync_audio", _speakerConfigParam4SyncAudio)
    return nil
}

// SpeakerConfigParam4SyncAudio Getter
func (r AlibabaLstSpeakerConfigureSyncaudioRequest) GetSpeakerConfigParam4SyncAudio() *SpeakerConfigParam4SyncAudio {
    return r._speakerConfigParam4SyncAudio
}
