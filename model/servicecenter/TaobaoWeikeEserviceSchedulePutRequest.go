package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交客服排班信息 API请求
taobao.weike.eservice.schedule.put

添加、更新、删除排班信息
*/
type TaobaoWeikeEserviceSchedulePutRequest struct {
    model.Params
    // 订单ID
    _orderId   int64
    // 按天排班信息
    _csSchedulings   []CsSchedulingOneDayDto
}

// 初始化TaobaoWeikeEserviceSchedulePutRequest对象
func NewTaobaoWeikeEserviceSchedulePutRequest() *TaobaoWeikeEserviceSchedulePutRequest{
    return &TaobaoWeikeEserviceSchedulePutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceSchedulePutRequest) GetApiMethodName() string {
    return "taobao.weike.eservice.schedule.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceSchedulePutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *TaobaoWeikeEserviceSchedulePutRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoWeikeEserviceSchedulePutRequest) GetOrderId() int64 {
    return r._orderId
}
// CsSchedulings Setter
// 按天排班信息
func (r *TaobaoWeikeEserviceSchedulePutRequest) SetCsSchedulings(_csSchedulings []CsSchedulingOneDayDto) error {
    r._csSchedulings = _csSchedulings
    r.Set("cs_schedulings", _csSchedulings)
    return nil
}

// CsSchedulings Getter
func (r TaobaoWeikeEserviceSchedulePutRequest) GetCsSchedulings() []CsSchedulingOneDayDto {
    return r._csSchedulings
}
