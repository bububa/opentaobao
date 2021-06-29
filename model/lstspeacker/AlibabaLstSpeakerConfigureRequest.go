package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通音箱配置通用泛化调用接口 APIRequest
alibaba.lst.speaker.configure

零售通音箱配置通用泛化调用接口，包括内容、音量、音频等内容
*/
type AlibabaLstSpeakerConfigureRequest struct {
    model.Params

    // 设备编码
    deviceCode   string 

    // 命令类型setPayTime,adjustVolume，syncAudio，syncAudioAdvert
    command   string 

    // 数据体，根据命令不同而不同
    params   string 

}

func NewAlibabaLstSpeakerConfigureRequest() *AlibabaLstSpeakerConfigureRequest{
    return &AlibabaLstSpeakerConfigureRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstSpeakerConfigureRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.configure"
}

func (r AlibabaLstSpeakerConfigureRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstSpeakerConfigureRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r AlibabaLstSpeakerConfigureRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *AlibabaLstSpeakerConfigureRequest) SetCommand(command string) error {
    r.command = command
    r.Set("command", command)
    return nil
}

func (r AlibabaLstSpeakerConfigureRequest) GetCommand() string {
    return r.command
}

func (r *AlibabaLstSpeakerConfigureRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaLstSpeakerConfigureRequest) GetParams() string {
    return r.params
}

