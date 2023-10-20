package tbitem

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemCombineGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
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

var poolTmallItemCombineGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemCombineGetRequest()
	},
}

// GetTmallItemCombineGetRequest 从 sync.Pool 获取 TmallItemCombineGetAPIRequest
func GetTmallItemCombineGetAPIRequest() *TmallItemCombineGetAPIRequest {
	return poolTmallItemCombineGetAPIRequest.Get().(*TmallItemCombineGetAPIRequest)
}

// ReleaseTmallItemCombineGetAPIRequest 将 TmallItemCombineGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemCombineGetAPIRequest(v *TmallItemCombineGetAPIRequest) {
	v.Reset()
	poolTmallItemCombineGetAPIRequest.Put(v)
}
