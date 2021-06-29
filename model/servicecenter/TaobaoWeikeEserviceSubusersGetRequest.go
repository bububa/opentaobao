package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客服外包订单分配的商家子账号列表 API请求
taobao.weike.eservice.subusers.get

获取客服外包订单分配的商家子账号列表，以及授权状态
*/
type TaobaoWeikeEserviceSubusersGetRequest struct {
    model.Params
    // 订单ID
    _orderId   int64
}

// 初始化TaobaoWeikeEserviceSubusersGetRequest对象
func NewTaobaoWeikeEserviceSubusersGetRequest() *TaobaoWeikeEserviceSubusersGetRequest{
    return &TaobaoWeikeEserviceSubusersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceSubusersGetRequest) GetApiMethodName() string {
    return "taobao.weike.eservice.subusers.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceSubusersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *TaobaoWeikeEserviceSubusersGetRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoWeikeEserviceSubusersGetRequest) GetOrderId() int64 {
    return r._orderId
}
