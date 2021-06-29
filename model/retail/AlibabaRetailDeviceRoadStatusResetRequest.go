package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机货道解锁 APIRequest
alibaba.retail.device.road.status.reset

贩卖机货道解锁
*/
type AlibabaRetailDeviceRoadStatusResetRequest struct {
    model.Params

    // 设备外部编码
    deviceUuid   string 

    // 阿里设备编码
    deviceCode   string 

    // 阿里设备物理编码
    deviceSn   string 

    // 货道编码
    roadNoList   []string 

}

func NewAlibabaRetailDeviceRoadStatusResetRequest() *AlibabaRetailDeviceRoadStatusResetRequest{
    return &AlibabaRetailDeviceRoadStatusResetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailDeviceRoadStatusResetRequest) GetApiMethodName() string {
    return "alibaba.retail.device.road.status.reset"
}

func (r AlibabaRetailDeviceRoadStatusResetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailDeviceRoadStatusResetRequest) SetDeviceUuid(deviceUuid string) error {
    r.deviceUuid = deviceUuid
    r.Set("device_uuid", deviceUuid)
    return nil
}

func (r AlibabaRetailDeviceRoadStatusResetRequest) GetDeviceUuid() string {
    return r.deviceUuid
}

func (r *AlibabaRetailDeviceRoadStatusResetRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r AlibabaRetailDeviceRoadStatusResetRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *AlibabaRetailDeviceRoadStatusResetRequest) SetDeviceSn(deviceSn string) error {
    r.deviceSn = deviceSn
    r.Set("device_sn", deviceSn)
    return nil
}

func (r AlibabaRetailDeviceRoadStatusResetRequest) GetDeviceSn() string {
    return r.deviceSn
}

func (r *AlibabaRetailDeviceRoadStatusResetRequest) SetRoadNoList(roadNoList []string) error {
    r.roadNoList = roadNoList
    r.Set("road_no_list", roadNoList)
    return nil
}

func (r AlibabaRetailDeviceRoadStatusResetRequest) GetRoadNoList() []string {
    return r.roadNoList
}

