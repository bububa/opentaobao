package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuSortGetAPIRequest sku销售属性顺序获取 API请求
// tmall.item.sku.sort.get
//
// sku销售属性顺序获取
type TmallItemSkuSortGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallItemSkuSortGetRequest 初始化TmallItemSkuSortGetAPIRequest对象
func NewTmallItemSkuSortGetRequest() *TmallItemSkuSortGetAPIRequest {
	return &TmallItemSkuSortGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSkuSortGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSkuSortGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.sort.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSkuSortGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSkuSortGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemSkuSortGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSkuSortGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallItemSkuSortGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSkuSortGetRequest()
	},
}

// GetTmallItemSkuSortGetRequest 从 sync.Pool 获取 TmallItemSkuSortGetAPIRequest
func GetTmallItemSkuSortGetAPIRequest() *TmallItemSkuSortGetAPIRequest {
	return poolTmallItemSkuSortGetAPIRequest.Get().(*TmallItemSkuSortGetAPIRequest)
}

// ReleaseTmallItemSkuSortGetAPIRequest 将 TmallItemSkuSortGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemSkuSortGetAPIRequest(v *TmallItemSkuSortGetAPIRequest) {
	v.Reset()
	poolTmallItemSkuSortGetAPIRequest.Put(v)
}
