package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalldeviceitempromotiongetAPIRequest 智能硬件上商品优惠获取 API请求
// tmall.device.item.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
type TmalldeviceitempromotiongetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmalldeviceitempromotiongetRequest 初始化TmalldeviceitempromotiongetAPIRequest对象
func NewTmalldeviceitempromotiongetRequest() *TmalldeviceitempromotiongetAPIRequest {
	return &TmalldeviceitempromotiongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalldeviceitempromotiongetAPIRequest) GetApiMethodName() string {
	return "tmall.device.item.promotion.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalldeviceitempromotiongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalldeviceitempromotiongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmalldeviceitempromotiongetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmalldeviceitempromotiongetAPIRequest) GetItemId() int64 {
	return r._itemId
}
