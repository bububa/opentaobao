package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通音箱配置通用泛化调用接口 API请求
alibaba.lst.speaker.configure

零售通音箱配置通用泛化调用接口，包括内容、音量、音频等内容
*/
type AlibabaLstSpeakerConfigureRequest struct {
    model.Params
    // 设备编码
    _deviceCode   string
    // 命令类型setPayTime,adjustVolume，syncAudio，syncAudioAdvert
    _command   string
    // 数据体，根据命令不同而不同
    _params   string
}

// 初始化AlibabaLstSpeakerConfigureRequest对象
func NewAlibabaLstSpeakerConfigureRequest() *AlibabaLstSpeakerConfigureRequest{
    return &AlibabaLstSpeakerConfigureRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerConfigureRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.configure"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerConfigureRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerConfigureRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaLstSpeakerConfigureRequest) GetDeviceCode() string {
    return r._deviceCode
}
// Command Setter
// 命令类型setPayTime,adjustVolume，syncAudio，syncAudioAdvert
func (r *AlibabaLstSpeakerConfigureRequest) SetCommand(_command string) error {
    r._command = _command
    r.Set("command", _command)
    return nil
}

// Command Getter
func (r AlibabaLstSpeakerConfigureRequest) GetCommand() string {
    return r._command
}
// Params Setter
// 数据体，根据命令不同而不同
func (r *AlibabaLstSpeakerConfigureRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaLstSpeakerConfigureRequest) GetParams() string {
    return r._params
}
