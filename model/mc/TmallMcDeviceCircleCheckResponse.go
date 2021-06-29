package mc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云码设备圈选情况查询 API返回值 
tmall.mc.device.circle.check

云码设备圈选情况查询
*/
type TmallMcDeviceCircleCheckAPIResponse struct {
    model.CommonResponse
    TmallMcDeviceCircleCheckResponse
}

// 云码设备圈选情况查询 成功返回结果
type TmallMcDeviceCircleCheckResponse struct {
    XMLName xml.Name `xml:"tmall_mc_device_circle_check_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 设备相关投放计划
    Results   []TaskDTO `json:"results,omitempty" xml:"results>task_dto,omitempty"`
}
