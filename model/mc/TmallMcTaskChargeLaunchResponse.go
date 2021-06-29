package mc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云码充电宝投放链路 API返回值 
tmall.mc.task.charge.launch

云码充电宝投放链路，用于判断用户是否有匹配的投放计划
*/
type TmallMcTaskChargeLaunchAPIResponse struct {
    model.CommonResponse
    TmallMcTaskChargeLaunchResponse
}

// 云码充电宝投放链路 成功返回结果
type TmallMcTaskChargeLaunchResponse struct {
    XMLName xml.Name `xml:"tmall_mc_task_charge_launch_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 入会页url
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
