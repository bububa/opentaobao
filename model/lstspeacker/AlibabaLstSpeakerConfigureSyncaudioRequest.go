package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
音频同步 APIRequest
alibaba.lst.speaker.configure.syncaudio

音频同步
*/
type AlibabaLstSpeakerConfigureSyncaudioRequest struct {
    model.Params

    // 设备编码
    deviceCode   string 

    // 参数
    speakerConfigParam4SyncAudio   *SpeakerConfigParam4SyncAudio 

}

func NewAlibabaLstSpeakerConfigureSyncaudioRequest() *AlibabaLstSpeakerConfigureSyncaudioRequest{
    return &AlibabaLstSpeakerConfigureSyncaudioRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstSpeakerConfigureSyncaudioRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.configure.syncaudio"
}

func (r AlibabaLstSpeakerConfigureSyncaudioRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstSpeakerConfigureSyncaudioRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r AlibabaLstSpeakerConfigureSyncaudioRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *AlibabaLstSpeakerConfigureSyncaudioRequest) SetSpeakerConfigParam4SyncAudio(speakerConfigParam4SyncAudio *SpeakerConfigParam4SyncAudio) error {
    r.speakerConfigParam4SyncAudio = speakerConfigParam4SyncAudio
    r.Set("speaker_config_param4_sync_audio", speakerConfigParam4SyncAudio)
    return nil
}

func (r AlibabaLstSpeakerConfigureSyncaudioRequest) GetSpeakerConfigParam4SyncAudio() *SpeakerConfigParam4SyncAudio {
    return r.speakerConfigParam4SyncAudio
}

