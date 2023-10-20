package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcombooffshelfAPIRequest 酒店套餐下架 API请求
// taobao.xhotel.combo.offshelf
//
// 酒店套餐下架
type TaobaoxhotelcombooffshelfAPIRequest struct {
	model.Params
	// 宝贝id
	_itemId int64
}

// NewTaobaoxhotelcombooffshelfRequest 初始化TaobaoxhotelcombooffshelfAPIRequest对象
func NewTaobaoxhotelcombooffshelfRequest() *TaobaoxhotelcombooffshelfAPIRequest {
	return &TaobaoxhotelcombooffshelfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelcombooffshelfAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.combo.offshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelcombooffshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelcombooffshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 宝贝id
func (r *TaobaoxhotelcombooffshelfAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoxhotelcombooffshelfAPIRequest) GetItemId() int64 {
	return r._itemId
}
