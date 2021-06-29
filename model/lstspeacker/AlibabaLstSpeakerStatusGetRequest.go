package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
音箱设备在线状态 APIRequest
alibaba.lst.speaker.status.get

音箱设备在线状态查询
*/
type AlibabaLstSpeakerStatusGetRequest struct {
    model.Params

    // 设备编码
    deviceCode   string 

}

func NewAlibabaLstSpeakerStatusGetRequest() *AlibabaLstSpeakerStatusGetRequest{
    return &AlibabaLstSpeakerStatusGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstSpeakerStatusGetRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.status.get"
}

func (r AlibabaLstSpeakerStatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstSpeakerStatusGetRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r AlibabaLstSpeakerStatusGetRequest) GetDeviceCode() string {
    return r.deviceCode
}

