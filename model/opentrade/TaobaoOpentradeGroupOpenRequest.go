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
type TaobaoOpentradeGroupOpenRequest struct {
    model.Params
    // 商品id
    itemId   int64
    // 用户openId
    openUserId   string
    // 组团活动id
    groupActivityId   int64
}

// 初始化TaobaoOpentradeGroupOpenRequest对象
func NewTaobaoOpentradeGroupOpenRequest() *TaobaoOpentradeGroupOpenRequest{
    return &TaobaoOpentradeGroupOpenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupOpenRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.open"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupOpenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoOpentradeGroupOpenRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupOpenRequest) GetItemId() int64 {
    return r.itemId
}
// OpenUserId Setter
// 用户openId
func (r *TaobaoOpentradeGroupOpenRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

// OpenUserId Getter
func (r TaobaoOpentradeGroupOpenRequest) GetOpenUserId() string {
    return r.openUserId
}
// GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupOpenRequest) SetGroupActivityId(groupActivityId int64) error {
    r.groupActivityId = groupActivityId
    r.Set("group_activity_id", groupActivityId)
    return nil
}

// GroupActivityId Getter
func (r TaobaoOpentradeGroupOpenRequest) GetGroupActivityId() int64 {
    return r.groupActivityId
}
