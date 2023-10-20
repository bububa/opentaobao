package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumppromotiongetAPIRequest 商品优惠详情查询 API请求
// taobao.ump.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
type TaobaoumppromotiongetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTaobaoumppromotiongetRequest 初始化TaobaoumppromotiongetAPIRequest对象
func NewTaobaoumppromotiongetRequest() *TaobaoumppromotiongetAPIRequest {
	return &TaobaoumppromotiongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumppromotiongetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.promotion.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumppromotiongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumppromotiongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoumppromotiongetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoumppromotiongetAPIRequest) GetItemId() int64 {
	return r._itemId
}
