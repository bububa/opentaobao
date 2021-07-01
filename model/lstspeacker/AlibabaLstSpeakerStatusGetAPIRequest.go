package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
音箱设备在线状态 API请求
alibaba.lst.speaker.status.get

音箱设备在线状态查询
*/
type AlibabaLstSpeakerStatusGetAPIRequest struct {
    model.Params
    // 设备编码
    _deviceCode   string
}

// 初始化AlibabaLstSpeakerStatusGetAPIRequest对象
func NewAlibabaLstSpeakerStatusGetRequest() *AlibabaLstSpeakerStatusGetAPIRequest{
    return &AlibabaLstSpeakerStatusGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerStatusGetAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.status.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerStatusGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerStatusGetAPIRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaLstSpeakerStatusGetAPIRequest) GetDeviceCode() string {
    return r._deviceCode
}
