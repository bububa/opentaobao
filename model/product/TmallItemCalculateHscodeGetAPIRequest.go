package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemCalculateHscodeGetAPIRequest 算法获取hscode API请求
// tmall.item.calculate.hscode.get
//
// 算法获取hscode
type TmallItemCalculateHscodeGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallItemCalculateHscodeGetRequest 初始化TmallItemCalculateHscodeGetAPIRequest对象
func NewTmallItemCalculateHscodeGetRequest() *TmallItemCalculateHscodeGetAPIRequest {
	return &TmallItemCalculateHscodeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemCalculateHscodeGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.calculate.hscode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemCalculateHscodeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallItemCalculateHscodeGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemCalculateHscodeGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
