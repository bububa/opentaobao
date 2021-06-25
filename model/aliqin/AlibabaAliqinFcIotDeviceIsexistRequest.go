package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断设备是否存在 APIRequest
alibaba.aliqin.fc.iot.device.isexist

判断设备是否存在
*/
type AlibabaAliqinFcIotDeviceIsexistRequest struct {
    model.Params

    // 设备编号
    imei   string 

    // 设备类型（预留将来扩展）
    deviceType   string 

    // 渠道扩展编码（预留）
    midPatChannel   string 

}

func NewAlibabaAliqinFcIotDeviceIsexistRequest() *AlibabaAliqinFcIotDeviceIsexistRequest{
    return &AlibabaAliqinFcIotDeviceIsexistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.device.isexist"
}

func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcIotDeviceIsexistRequest) SetImei(imei string) error {
    r.imei = imei
    r.Set("imei", imei)
    return nil
}

func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetImei() string {
    return r.imei
}

func (r *AlibabaAliqinFcIotDeviceIsexistRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetDeviceType() string {
    return r.deviceType
}

func (r *AlibabaAliqinFcIotDeviceIsexistRequest) SetMidPatChannel(midPatChannel string) error {
    r.midPatChannel = midPatChannel
    r.Set("mid_pat_channel", midPatChannel)
    return nil
}

func (r AlibabaAliqinFcIotDeviceIsexistRequest) GetMidPatChannel() string {
    return r.midPatChannel
}

