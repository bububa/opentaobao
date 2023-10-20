package omniorder

import (
	"net/url"
	"sync"

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
	// 商品outerId
	_outerId string
	// 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
	_itemId int64
}

// NewTaobaoOmniitemItemDeleteRequest 初始化TaobaoOmniitemItemDeleteAPIRequest对象
func NewTaobaoOmniitemItemDeleteRequest() *TaobaoOmniitemItemDeleteAPIRequest {
	return &TaobaoOmniitemItemDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniitemItemDeleteAPIRequest) Reset() {
	r._barCode = ""
	r._outerId = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniitemItemDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBarCode is BarCode Setter
// 条形码
func (r *TaobaoOmniitemItemDeleteAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r TaobaoOmniitemItemDeleteAPIRequest) GetBarCode() string {
	return r._barCode
}

// SetOuterId is OuterId Setter
// 商品outerId
func (r *TaobaoOmniitemItemDeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoOmniitemItemDeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetItemId is ItemId Setter
// 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
func (r *TaobaoOmniitemItemDeleteAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoOmniitemItemDeleteAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoOmniitemItemDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniitemItemDeleteRequest()
	},
}

// GetTaobaoOmniitemItemDeleteRequest 从 sync.Pool 获取 TaobaoOmniitemItemDeleteAPIRequest
func GetTaobaoOmniitemItemDeleteAPIRequest() *TaobaoOmniitemItemDeleteAPIRequest {
	return poolTaobaoOmniitemItemDeleteAPIRequest.Get().(*TaobaoOmniitemItemDeleteAPIRequest)
}

// ReleaseTaobaoOmniitemItemDeleteAPIRequest 将 TaobaoOmniitemItemDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniitemItemDeleteAPIRequest(v *TaobaoOmniitemItemDeleteAPIRequest) {
	v.Reset()
	poolTaobaoOmniitemItemDeleteAPIRequest.Put(v)
}
