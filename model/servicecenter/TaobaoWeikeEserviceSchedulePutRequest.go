package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交客服排班信息 APIRequest
taobao.weike.eservice.schedule.put

添加、更新、删除排班信息
*/
type TaobaoWeikeEserviceSchedulePutRequest struct {
    model.Params

    // 订单ID
    orderId   int64 

    // 按天排班信息
    csSchedulings   []CsSchedulingOneDayDto 

}

func NewTaobaoWeikeEserviceSchedulePutRequest() *TaobaoWeikeEserviceSchedulePutRequest{
    return &TaobaoWeikeEserviceSchedulePutRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWeikeEserviceSchedulePutRequest) GetApiMethodName() string {
    return "taobao.weike.eservice.schedule.put"
}

func (r TaobaoWeikeEserviceSchedulePutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWeikeEserviceSchedulePutRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoWeikeEserviceSchedulePutRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoWeikeEserviceSchedulePutRequest) SetCsSchedulings(csSchedulings []CsSchedulingOneDayDto) error {
    r.csSchedulings = csSchedulings
    r.Set("cs_schedulings", csSchedulings)
    return nil
}

func (r TaobaoWeikeEserviceSchedulePutRequest) GetCsSchedulings() []CsSchedulingOneDayDto {
    return r.csSchedulings
}

