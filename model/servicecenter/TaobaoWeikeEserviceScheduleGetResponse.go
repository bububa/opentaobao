package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
客服排班信息查询接口 APIResponse
taobao.weike.eservice.schedule.get

客服排班信息查询接口
*/
type TaobaoWeikeEserviceScheduleGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWeikeEserviceScheduleGetResponse `json:"weike_eservice_schedule_get_response,omitempty"` 
    TaobaoWeikeEserviceScheduleGetResponse
}

/* model for simplify = false
type TaobaoWeikeEserviceScheduleGetResponse struct {

    // 排班信息查询结果
    
    Result  *struct {
        CsSchedulingWrapper  *CsSchedulingWrapper `json:"cs_scheduling_wrapper,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWeikeEserviceScheduleGetResponse struct {

    // 排班信息查询结果
    Result   *CsSchedulingWrapper `json:"result,omitempty"`

}
