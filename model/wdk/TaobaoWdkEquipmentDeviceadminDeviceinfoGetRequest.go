package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取五道口设备管理信息 APIRequest
taobao.wdk.equipment.deviceadmin.deviceinfo.get

通过仓编码获取五道口设备管理信息
*/
type TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest struct {
    model.Params

    // 仓编码
    warehouseCode   string 

    // 设备类型
    deviceType   int64 

}

func NewTaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest() *TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest{
    return &TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.deviceadmin.deviceinfo.get"
}

func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

func (r *TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) GetDeviceType() int64 {
    return r.deviceType
}

