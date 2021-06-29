package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机设备信息获取 APIRequest
alibaba.retail.device.info.get

贩卖机设备信息获取
*/
type AlibabaRetailDeviceInfoGetRequest struct {
    model.Params

    // 外部设备ID
    deviceUuid   string 

    // 阿里设备编码
    deviceCode   string 

    // 阿里设备物理ID（32位）
    deviceSn   string 

}

func NewAlibabaRetailDeviceInfoGetRequest() *AlibabaRetailDeviceInfoGetRequest{
    return &AlibabaRetailDeviceInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailDeviceInfoGetRequest) GetApiMethodName() string {
    return "alibaba.retail.device.info.get"
}

func (r AlibabaRetailDeviceInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailDeviceInfoGetRequest) SetDeviceUuid(deviceUuid string) error {
    r.deviceUuid = deviceUuid
    r.Set("device_uuid", deviceUuid)
    return nil
}

func (r AlibabaRetailDeviceInfoGetRequest) GetDeviceUuid() string {
    return r.deviceUuid
}

func (r *AlibabaRetailDeviceInfoGetRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r AlibabaRetailDeviceInfoGetRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *AlibabaRetailDeviceInfoGetRequest) SetDeviceSn(deviceSn string) error {
    r.deviceSn = deviceSn
    r.Set("device_sn", deviceSn)
    return nil
}

func (r AlibabaRetailDeviceInfoGetRequest) GetDeviceSn() string {
    return r.deviceSn
}

