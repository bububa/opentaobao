package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelComboOffshelfAPIRequest 酒店套餐下架 API请求
// taobao.xhotel.combo.offshelf
//
// 酒店套餐下架
type TaobaoXhotelComboOffshelfAPIRequest struct {
	model.Params
	// 宝贝id
	_itemId int64
}

// NewTaobaoXhotelComboOffshelfRequest 初始化TaobaoXhotelComboOffshelfAPIRequest对象
func NewTaobaoXhotelComboOffshelfRequest() *TaobaoXhotelComboOffshelfAPIRequest {
	return &TaobaoXhotelComboOffshelfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelComboOffshelfAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.combo.offshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelComboOffshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelComboOffshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 宝贝id
func (r *TaobaoXhotelComboOffshelfAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoXhotelComboOffshelfAPIRequest) GetItemId() int64 {
	return r._itemId
}
