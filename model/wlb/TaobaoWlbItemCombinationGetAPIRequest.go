package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbitemcombinationgetAPIRequest 根据商品id查询商品组合关系 API请求
// taobao.wlb.item.combination.get
//
// 根据商品id查询商品组合关系
type TaobaowlbitemcombinationgetAPIRequest struct {
	model.Params
	// 要查询的组合商品id
	_itemId int64
}

// NewTaobaowlbitemcombinationgetRequest 初始化TaobaowlbitemcombinationgetAPIRequest对象
func NewTaobaowlbitemcombinationgetRequest() *TaobaowlbitemcombinationgetAPIRequest {
	return &TaobaowlbitemcombinationgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbitemcombinationgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.combination.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbitemcombinationgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbitemcombinationgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 要查询的组合商品id
func (r *TaobaowlbitemcombinationgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaowlbitemcombinationgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
