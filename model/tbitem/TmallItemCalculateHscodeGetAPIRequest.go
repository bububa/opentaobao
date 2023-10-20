package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemcalculatehscodegetAPIRequest 算法获取hscode API请求
// tmall.item.calculate.hscode.get
//
// 算法获取hscode
type TmallitemcalculatehscodegetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallitemcalculatehscodegetRequest 初始化TmallitemcalculatehscodegetAPIRequest对象
func NewTmallitemcalculatehscodegetRequest() *TmallitemcalculatehscodegetAPIRequest {
	return &TmallitemcalculatehscodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemcalculatehscodegetAPIRequest) GetApiMethodName() string {
	return "tmall.item.calculate.hscode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemcalculatehscodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemcalculatehscodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallitemcalculatehscodegetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemcalculatehscodegetAPIRequest) GetItemId() int64 {
	return r._itemId
}
