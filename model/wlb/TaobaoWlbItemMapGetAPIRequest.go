package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbitemmapgetAPIRequest 根据物流宝商品ID查询商品映射关系 API请求
// taobao.wlb.item.map.get
//
// 根据物流宝商品ID查询商品映射关系
type TaobaowlbitemmapgetAPIRequest struct {
	model.Params
	// 要查询映射关系的物流宝商品id
	_itemId int64
}

// NewTaobaowlbitemmapgetRequest 初始化TaobaowlbitemmapgetAPIRequest对象
func NewTaobaowlbitemmapgetRequest() *TaobaowlbitemmapgetAPIRequest {
	return &TaobaowlbitemmapgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbitemmapgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.map.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbitemmapgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbitemmapgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 要查询映射关系的物流宝商品id
func (r *TaobaowlbitemmapgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaowlbitemmapgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
