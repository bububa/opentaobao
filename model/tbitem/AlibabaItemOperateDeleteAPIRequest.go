package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitemoperatedeleteAPIRequest 商品删除 API请求
// alibaba.item.operate.delete
//
// 商品删除
type AlibabaitemoperatedeleteAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewAlibabaitemoperatedeleteRequest 初始化AlibabaitemoperatedeleteAPIRequest对象
func NewAlibabaitemoperatedeleteRequest() *AlibabaitemoperatedeleteAPIRequest {
	return &AlibabaitemoperatedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitemoperatedeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.item.operate.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitemoperatedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitemoperatedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaitemoperatedeleteAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaitemoperatedeleteAPIRequest) GetItemId() int64 {
	return r._itemId
}
