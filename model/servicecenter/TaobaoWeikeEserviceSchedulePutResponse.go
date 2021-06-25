package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
提交客服排班信息 APIResponse
taobao.weike.eservice.schedule.put

添加、更新、删除排班信息
*/
type TaobaoWeikeEserviceSchedulePutAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWeikeEserviceSchedulePutResponse `json:"taobao_weike_eservice_schedule_put_response,omitempty"`
}

type TaobaoWeikeEserviceSchedulePutResponse struct {

    // 是否执行成功
    Result   bool `json:"result,omitempty"`

}
