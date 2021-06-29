package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
音箱音量调节 API请求
alibaba.lst.speaker.configure.adjustvolume

音箱音量调节
*/
type AlibabaLstSpeakerConfigureAdjustvolumeRequest struct {
    model.Params
    // 设备编码
    deviceCode   string
    // 音量直
    volume   string
    // 音量类型，val:固定值, percent:百分比
    valueType   string
}

// 初始化AlibabaLstSpeakerConfigureAdjustvolumeRequest对象
func NewAlibabaLstSpeakerConfigureAdjustvolumeRequest() *AlibabaLstSpeakerConfigureAdjustvolumeRequest{
    return &AlibabaLstSpeakerConfigureAdjustvolumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerConfigureAdjustvolumeRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.configure.adjustvolume"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerConfigureAdjustvolumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerConfigureAdjustvolumeRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaLstSpeakerConfigureAdjustvolumeRequest) GetDeviceCode() string {
    return r.deviceCode
}
// Volume Setter
// 音量直
func (r *AlibabaLstSpeakerConfigureAdjustvolumeRequest) SetVolume(volume string) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

// Volume Getter
func (r AlibabaLstSpeakerConfigureAdjustvolumeRequest) GetVolume() string {
    return r.volume
}
// ValueType Setter
// 音量类型，val:固定值, percent:百分比
func (r *AlibabaLstSpeakerConfigureAdjustvolumeRequest) SetValueType(valueType string) error {
    r.valueType = valueType
    r.Set("value_type", valueType)
    return nil
}

// ValueType Getter
func (r AlibabaLstSpeakerConfigureAdjustvolumeRequest) GetValueType() string {
    return r.valueType
}
