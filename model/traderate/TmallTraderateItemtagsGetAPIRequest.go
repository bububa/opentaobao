package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltraderateitemtagsgetAPIRequest 通过商品ID获取标签列表 API请求
// tmall.traderate.itemtags.get
//
// 通过商品ID获取标签详细信息
type TmalltraderateitemtagsgetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmalltraderateitemtagsgetRequest 初始化TmalltraderateitemtagsgetAPIRequest对象
func NewTmalltraderateitemtagsgetRequest() *TmalltraderateitemtagsgetAPIRequest {
	return &TmalltraderateitemtagsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltraderateitemtagsgetAPIRequest) GetApiMethodName() string {
	return "tmall.traderate.itemtags.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltraderateitemtagsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltraderateitemtagsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmalltraderateitemtagsgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmalltraderateitemtagsgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
