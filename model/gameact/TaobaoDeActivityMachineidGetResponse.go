package gameact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备号 APIResponse
taobao.de.activity.machineid.get

获取机器设备id
*/
type TaobaoDeActivityMachineidGetAPIResponse struct {
    model.CommonResponse
    TaobaoDeActivityMachineidGetResponse
}

type TaobaoDeActivityMachineidGetResponse struct {
    XMLName xml.Name `xml:"de_activity_machineid_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 机器号
    
    MachineId   string `json:"machine_id,omitempty" xml:"machine_id,omitempty"`

    
}
