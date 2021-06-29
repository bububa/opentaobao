package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购获取订单列表 API请求
taobao.opentrade.group.order

组团购场景下，获取开团的订单列表
*/
type TaobaoOpentradeGroupOrderRequest struct {
    model.Params
    // 团id
    _groupId   int64
}

// 初始化TaobaoOpentradeGroupOrderRequest对象
func NewTaobaoOpentradeGroupOrderRequest() *TaobaoOpentradeGroupOrderRequest{
    return &TaobaoOpentradeGroupOrderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupOrderRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.order"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupOrderRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r TaobaoOpentradeGroupOrderRequest) GetGroupId() int64 {
    return r._groupId
}
