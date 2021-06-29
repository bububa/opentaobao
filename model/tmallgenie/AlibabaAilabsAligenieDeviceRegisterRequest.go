package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵开放平台获取设备秘钥 APIRequest
alibaba.ailabs.aligenie.device.register

向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥
*/
type AlibabaAilabsAligenieDeviceRegisterRequest struct {
    model.Params

    // 设备id
    deviceId   int64 

    // mac区段脚本
    macSections   string 

}

func NewAlibabaAilabsAligenieDeviceRegisterRequest() *AlibabaAilabsAligenieDeviceRegisterRequest{
    return &AlibabaAilabsAligenieDeviceRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieDeviceRegisterRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.device.register"
}

func (r AlibabaAilabsAligenieDeviceRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieDeviceRegisterRequest) SetDeviceId(deviceId int64) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r AlibabaAilabsAligenieDeviceRegisterRequest) GetDeviceId() int64 {
    return r.deviceId
}

func (r *AlibabaAilabsAligenieDeviceRegisterRequest) SetMacSections(macSections string) error {
    r.macSections = macSections
    r.Set("mac_sections", macSections)
    return nil
}

func (r AlibabaAilabsAligenieDeviceRegisterRequest) GetMacSections() string {
    return r.macSections
}

