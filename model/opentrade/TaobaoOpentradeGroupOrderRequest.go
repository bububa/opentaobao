package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购获取订单列表 APIRequest
taobao.opentrade.group.order

组团购场景下，获取开团的订单列表
*/
type TaobaoOpentradeGroupOrderRequest struct {
    model.Params

    // 团id
    groupId   int64 

}

func NewTaobaoOpentradeGroupOrderRequest() *TaobaoOpentradeGroupOrderRequest{
    return &TaobaoOpentradeGroupOrderRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeGroupOrderRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.order"
}

func (r TaobaoOpentradeGroupOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeGroupOrderRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r TaobaoOpentradeGroupOrderRequest) GetGroupId() int64 {
    return r.groupId
}

