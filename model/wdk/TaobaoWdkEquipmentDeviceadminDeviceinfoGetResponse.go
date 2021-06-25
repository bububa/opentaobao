package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取五道口设备管理信息 APIResponse
taobao.wdk.equipment.deviceadmin.deviceinfo.get

通过仓编码获取五道口设备管理信息
*/
type TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWdkEquipmentDeviceadminDeviceinfoGetResponse `json:"taobao_wdk_equipment_deviceadmin_deviceinfo_get_response,omitempty"`
}

type TaobaoWdkEquipmentDeviceadminDeviceinfoGetResponse struct {

    // 返回值
    Result   *HmResult `json:"result,omitempty"`

}
