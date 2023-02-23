package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemCombineGetAPIRequest 组合商品获取接口 API请求
// tmall.item.combine.get
//
// 查询组合商品的SKU信息
type TmallItemCombineGetAPIRequest struct {
	model.Params
	// 组合商品ID
	_itemId int64
}

// NewTmallItemCombineGetRequest 初始化TmallItemCombineGetAPIRequest对象
func NewTmallItemCombineGetRequest() *TmallItemCombineGetAPIRequest {
	return &TmallItemCombineGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemCombineGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.combine.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemCombineGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemCombineGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 组合商品ID
func (r *TmallItemCombineGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemCombineGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
