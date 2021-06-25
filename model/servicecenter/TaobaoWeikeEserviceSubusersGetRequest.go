package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客服外包订单分配的商家子账号列表 APIRequest
taobao.weike.eservice.subusers.get

获取客服外包订单分配的商家子账号列表，以及授权状态
*/
type TaobaoWeikeEserviceSubusersGetRequest struct {
    model.Params

    // 订单ID
    orderId   int64 

}

func NewTaobaoWeikeEserviceSubusersGetRequest() *TaobaoWeikeEserviceSubusersGetRequest{
    return &TaobaoWeikeEserviceSubusersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWeikeEserviceSubusersGetRequest) GetApiMethodName() string {
    return "taobao.weike.eservice.subusers.get"
}

func (r TaobaoWeikeEserviceSubusersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWeikeEserviceSubusersGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoWeikeEserviceSubusersGetRequest) GetOrderId() int64 {
    return r.orderId
}

