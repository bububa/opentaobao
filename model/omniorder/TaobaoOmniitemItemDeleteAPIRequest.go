package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemItemDeleteAPIRequest 全渠道商品删除 API请求
// taobao.omniitem.item.delete
//
// 全渠道商品删除，能够对门店商品库商品进行删除动作
type TaobaoOmniitemItemDeleteAPIRequest struct {
	model.Params
	// 条形码
	_barCode string
	// 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
	_itemId int64
	// 商品outerId
	_outerId string
}

// NewTaobaoOmniitemItemDeleteRequest 初始化TaobaoOmniitemItemDeleteAPIRequest对象
func NewTaobaoOmniitemItemDeleteRequest() *TaobaoOmniitemItemDeleteAPIRequest {
	return &TaobaoOmniitemItemDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BarCode Setter
// 条形码
func (r *TaobaoOmniitemItemDeleteAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// Get BarCode Getter
func (r TaobaoOmniitemItemDeleteAPIRequest) GetBarCode() string {
	return r._barCode
}

// Set is ItemId Setter
// 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
func (r *TaobaoOmniitemItemDeleteAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoOmniitemItemDeleteAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is OuterId Setter
// 商品outerId
func (r *TaobaoOmniitemItemDeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoOmniitemItemDeleteAPIRequest) GetOuterId() string {
	return r._outerId
}
