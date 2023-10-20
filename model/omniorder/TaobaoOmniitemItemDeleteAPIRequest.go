package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniitemitemdeleteAPIRequest 全渠道商品删除 API请求
// taobao.omniitem.item.delete
//
// 全渠道商品删除，能够对门店商品库商品进行删除动作
type TaobaoomniitemitemdeleteAPIRequest struct {
	model.Params
	// 条形码
	_barCode string
	// 商品outerId
	_outerId string
	// 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
	_itemId int64
}

// NewTaobaoomniitemitemdeleteRequest 初始化TaobaoomniitemitemdeleteAPIRequest对象
func NewTaobaoomniitemitemdeleteRequest() *TaobaoomniitemitemdeleteAPIRequest {
	return &TaobaoomniitemitemdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniitemitemdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniitemitemdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniitemitemdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBarCode is BarCode Setter
// 条形码
func (r *TaobaoomniitemitemdeleteAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r TaobaoomniitemitemdeleteAPIRequest) GetBarCode() string {
	return r._barCode
}

// SetOuterId is OuterId Setter
// 商品outerId
func (r *TaobaoomniitemitemdeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoomniitemitemdeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetItemId is ItemId Setter
// 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
func (r *TaobaoomniitemitemdeleteAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoomniitemitemdeleteAPIRequest) GetItemId() int64 {
	return r._itemId
}
