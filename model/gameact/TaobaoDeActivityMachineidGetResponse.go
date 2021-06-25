package gameact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取设备号 APIResponse
taobao.de.activity.machineid.get

获取机器设备id
*/
type TaobaoDeActivityMachineidGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDeActivityMachineidGetResponse `json:"taobao_de_activity_machineid_get_response,omitempty"`
}

type TaobaoDeActivityMachineidGetResponse struct {

    // 机器号
    MachineId   string `json:"machine_id,omitempty"`

}
