package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemUnsubscribeAPIRequest 单个删除订阅关系 API请求
// taobao.baichuan.item.unsubscribe
//
// 删除单个商品订阅关系
type TaobaoBaichuanItemUnsubscribeAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTaobaoBaichuanItemUnsubscribeRequest 初始化TaobaoBaichuanItemUnsubscribeAPIRequest对象
func NewTaobaoBaichuanItemUnsubscribeRequest() *TaobaoBaichuanItemUnsubscribeAPIRequest {
	return &TaobaoBaichuanItemUnsubscribeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemUnsubscribeAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.item.unsubscribe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemUnsubscribeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanItemUnsubscribeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoBaichuanItemUnsubscribeAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoBaichuanItemUnsubscribeAPIRequest) GetItemId() int64 {
	return r._itemId
}
