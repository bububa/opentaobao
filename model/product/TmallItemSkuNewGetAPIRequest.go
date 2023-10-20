package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuNewGetAPIRequest 查询sku销售属性标新信息 API请求
// tmall.item.sku.new.get
//
// 查询sku销售属性标新信息
type TmallItemSkuNewGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallItemSkuNewGetRequest 初始化TmallItemSkuNewGetAPIRequest对象
func NewTmallItemSkuNewGetRequest() *TmallItemSkuNewGetAPIRequest {
	return &TmallItemSkuNewGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSkuNewGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSkuNewGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.new.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSkuNewGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSkuNewGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemSkuNewGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSkuNewGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallItemSkuNewGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSkuNewGetRequest()
	},
}

// GetTmallItemSkuNewGetRequest 从 sync.Pool 获取 TmallItemSkuNewGetAPIRequest
func GetTmallItemSkuNewGetAPIRequest() *TmallItemSkuNewGetAPIRequest {
	return poolTmallItemSkuNewGetAPIRequest.Get().(*TmallItemSkuNewGetAPIRequest)
}

// ReleaseTmallItemSkuNewGetAPIRequest 将 TmallItemSkuNewGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemSkuNewGetAPIRequest(v *TmallItemSkuNewGetAPIRequest) {
	v.Reset()
	poolTmallItemSkuNewGetAPIRequest.Put(v)
}
