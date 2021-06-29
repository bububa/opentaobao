package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景参团 API请求
taobao.opentrade.group.join

组团购场景下，用户参团
*/
type TaobaoOpentradeGroupJoinRequest struct {
    model.Params
    // 商品id
    itemId   int64
    // 团id
    groupId   int64
    // 用户openId
    openUserId   string
}

// 初始化TaobaoOpentradeGroupJoinRequest对象
func NewTaobaoOpentradeGroupJoinRequest() *TaobaoOpentradeGroupJoinRequest{
    return &TaobaoOpentradeGroupJoinRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupJoinRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.join"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupJoinRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoOpentradeGroupJoinRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupJoinRequest) GetItemId() int64 {
    return r.itemId
}
// GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupJoinRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r TaobaoOpentradeGroupJoinRequest) GetGroupId() int64 {
    return r.groupId
}
// OpenUserId Setter
// 用户openId
func (r *TaobaoOpentradeGroupJoinRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

// OpenUserId Getter
func (r TaobaoOpentradeGroupJoinRequest) GetOpenUserId() string {
    return r.openUserId
}
