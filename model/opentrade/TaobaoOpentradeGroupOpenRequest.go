package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景开团 API请求
taobao.opentrade.group.open

组团购场景下，团长开团
*/
type TaobaoOpentradeGroupOpenAPIRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // 用户openId
    _openUserId   string
    // 组团活动id
    _groupActivityId   int64
}

// 初始化TaobaoOpentradeGroupOpenAPIRequest对象
func NewTaobaoOpentradeGroupOpenRequest() *TaobaoOpentradeGroupOpenAPIRequest{
    return &TaobaoOpentradeGroupOpenAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupOpenAPIRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.open"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupOpenAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoOpentradeGroupOpenAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupOpenAPIRequest) GetItemId() int64 {
    return r._itemId
}
// OpenUserId Setter
// 用户openId
func (r *TaobaoOpentradeGroupOpenAPIRequest) SetOpenUserId(_openUserId string) error {
    r._openUserId = _openUserId
    r.Set("open_user_id", _openUserId)
    return nil
}

// OpenUserId Getter
func (r TaobaoOpentradeGroupOpenAPIRequest) GetOpenUserId() string {
    return r._openUserId
}
// GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupOpenAPIRequest) SetGroupActivityId(_groupActivityId int64) error {
    r._groupActivityId = _groupActivityId
    r.Set("group_activity_id", _groupActivityId)
    return nil
}

// GroupActivityId Getter
func (r TaobaoOpentradeGroupOpenAPIRequest) GetGroupActivityId() int64 {
    return r._groupActivityId
}
