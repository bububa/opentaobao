package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuStatusGetAPIRequest 商品sku上下架查询 API请求
// tmall.item.sku.status.get
//
// 商品sku上下架状态查询
type TmallItemSkuStatusGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallItemSkuStatusGetRequest 初始化TmallItemSkuStatusGetAPIRequest对象
func NewTmallItemSkuStatusGetRequest() *TmallItemSkuStatusGetAPIRequest {
	return &TmallItemSkuStatusGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSkuStatusGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSkuStatusGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSkuStatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSkuStatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemSkuStatusGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSkuStatusGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallItemSkuStatusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSkuStatusGetRequest()
	},
}

// GetTmallItemSkuStatusGetRequest 从 sync.Pool 获取 TmallItemSkuStatusGetAPIRequest
func GetTmallItemSkuStatusGetAPIRequest() *TmallItemSkuStatusGetAPIRequest {
	return poolTmallItemSkuStatusGetAPIRequest.Get().(*TmallItemSkuStatusGetAPIRequest)
}

// ReleaseTmallItemSkuStatusGetAPIRequest 将 TmallItemSkuStatusGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemSkuStatusGetAPIRequest(v *TmallItemSkuStatusGetAPIRequest) {
	v.Reset()
	poolTmallItemSkuStatusGetAPIRequest.Put(v)
}
