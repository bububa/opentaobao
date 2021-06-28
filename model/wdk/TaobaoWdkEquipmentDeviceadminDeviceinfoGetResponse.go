package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取五道口设备管理信息 APIResponse
taobao.wdk.equipment.deviceadmin.deviceinfo.get

通过仓编码获取五道口设备管理信息
*/
type TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentDeviceadminDeviceinfoGetResponse
}

type TaobaoWdkEquipmentDeviceadminDeviceinfoGetResponse struct {
    XMLName xml.Name `xml:"wdk_equipment_deviceadmin_deviceinfo_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *HmResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
