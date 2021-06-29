package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取五道口设备管理信息 API请求
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

// 初始化TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest对象
func NewTaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest() *TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest{
    return &TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.deviceadmin.deviceinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 仓编码
func (r *TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
// DeviceType Setter
// 设备类型
func (r *TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest) GetDeviceType() int64 {
    return r.deviceType
}
