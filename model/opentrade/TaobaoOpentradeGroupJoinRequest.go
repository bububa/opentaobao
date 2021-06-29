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
    _itemId   int64
    // 团id
    _groupId   int64
    // 用户openId
    _openUserId   string
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
func (r *TaobaoOpentradeGroupJoinRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupJoinRequest) GetItemId() int64 {
    return r._itemId
}
// GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupJoinRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r TaobaoOpentradeGroupJoinRequest) GetGroupId() int64 {
    return r._groupId
}
// OpenUserId Setter
// 用户openId
func (r *TaobaoOpentradeGroupJoinRequest) SetOpenUserId(_openUserId string) error {
    r._openUserId = _openUserId
    r.Set("open_user_id", _openUserId)
    return nil
}

// OpenUserId Getter
func (r TaobaoOpentradeGroupJoinRequest) GetOpenUserId() string {
    return r._openUserId
}
