package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeGroupJoinAPIRequest 组团购场景参团 API请求
// taobao.opentrade.group.join
//
// 组团购场景下，用户参团
type TaobaoOpentradeGroupJoinAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 团id
	_groupId int64
	// 用户openId
	_openUserId string
}

// NewTaobaoOpentradeGroupJoinRequest 初始化TaobaoOpentradeGroupJoinAPIRequest对象
func NewTaobaoOpentradeGroupJoinRequest() *TaobaoOpentradeGroupJoinAPIRequest {
	return &TaobaoOpentradeGroupJoinAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupJoinAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.group.join"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupJoinAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoOpentradeGroupJoinAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoOpentradeGroupJoinAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetGroupId is GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupJoinAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaoOpentradeGroupJoinAPIRequest) GetGroupId() int64 {
	return r._groupId
}

// SetOpenUserId is OpenUserId Setter
// 用户openId
func (r *TaobaoOpentradeGroupJoinAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// GetOpenUserId OpenUserId Getter
func (r TaobaoOpentradeGroupJoinAPIRequest) GetOpenUserId() string {
	return r._openUserId
}
