package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcomboreviewAPIRequest 套餐审核接口 API请求
// taobao.xhotel.combo.review
//
// 套餐审核接口
type TaobaoxhotelcomboreviewAPIRequest struct {
	model.Params
	// 宝贝id
	_itemId int64
}

// NewTaobaoxhotelcomboreviewRequest 初始化TaobaoxhotelcomboreviewAPIRequest对象
func NewTaobaoxhotelcomboreviewRequest() *TaobaoxhotelcomboreviewAPIRequest {
	return &TaobaoxhotelcomboreviewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelcomboreviewAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.combo.review"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelcomboreviewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelcomboreviewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 宝贝id
func (r *TaobaoxhotelcomboreviewAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoxhotelcomboreviewAPIRequest) GetItemId() int64 {
	return r._itemId
}
