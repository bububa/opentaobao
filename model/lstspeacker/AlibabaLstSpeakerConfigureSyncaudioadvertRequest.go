package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步广告 API请求
alibaba.lst.speaker.configure.syncaudioadvert

如意音箱广告同步
*/
type AlibabaLstSpeakerConfigureSyncaudioadvertRequest struct {
    model.Params
    // 设备编码
    _deviceCode   string
    // 音频参数
    _speakerConfigParam4SyncAudioAdvert   *SpeakerConfigParam4SyncAudioAdvert
}

// 初始化AlibabaLstSpeakerConfigureSyncaudioadvertRequest对象
func NewAlibabaLstSpeakerConfigureSyncaudioadvertRequest() *AlibabaLstSpeakerConfigureSyncaudioadvertRequest{
    return &AlibabaLstSpeakerConfigureSyncaudioadvertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerConfigureSyncaudioadvertRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.configure.syncaudioadvert"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerConfigureSyncaudioadvertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerConfigureSyncaudioadvertRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaLstSpeakerConfigureSyncaudioadvertRequest) GetDeviceCode() string {
    return r._deviceCode
}
// SpeakerConfigParam4SyncAudioAdvert Setter
// 音频参数
func (r *AlibabaLstSpeakerConfigureSyncaudioadvertRequest) SetSpeakerConfigParam4SyncAudioAdvert(_speakerConfigParam4SyncAudioAdvert *SpeakerConfigParam4SyncAudioAdvert) error {
    r._speakerConfigParam4SyncAudioAdvert = _speakerConfigParam4SyncAudioAdvert
    r.Set("speaker_config_param4_sync_audio_advert", _speakerConfigParam4SyncAudioAdvert)
    return nil
}

// SpeakerConfigParam4SyncAudioAdvert Getter
func (r AlibabaLstSpeakerConfigureSyncaudioadvertRequest) GetSpeakerConfigParam4SyncAudioAdvert() *SpeakerConfigParam4SyncAudioAdvert {
    return r._speakerConfigParam4SyncAudioAdvert
}
