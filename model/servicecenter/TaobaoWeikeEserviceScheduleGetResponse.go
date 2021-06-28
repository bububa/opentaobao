package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客服排班信息查询接口 APIResponse
taobao.weike.eservice.schedule.get

客服排班信息查询接口
*/
type TaobaoWeikeEserviceScheduleGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"weike_eservice_schedule_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 排班信息查询结果
    
    Result   *CsSchedulingWrapper `json:"result,omitempty" xml:"