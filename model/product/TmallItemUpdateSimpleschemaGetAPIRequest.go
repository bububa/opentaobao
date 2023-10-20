package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemUpdateSimpleschemaGetAPIRequest 官网同购编辑商品的get接口 API请求
// tmall.item.update.simpleschema.get
//
// 官网同购编辑商品的get接口
type TmallItemUpdateSimpleschemaGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallItemUpdateSimpleschemaGetRequest 初始化TmallItemUpdateSimpleschemaGetAPIRequest对象
func NewTmallItemUpdateSimpleschemaGetRequest() *TmallItemUpdateSimpleschemaGetAPIRequest {
	return &TmallItemUpdateSimpleschemaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemUpdateSimpleschemaGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemUpdateSimpleschemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.update.simpleschema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemUpdateSimpleschemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemUpdateSimpleschemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallItemUpdateSimpleschemaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemUpdateSimpleschemaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallItemUpdateSimpleschemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemUpdateSimpleschemaGetRequest()
	},
}

// GetTmallItemUpdateSimpleschemaGetRequest 从 sync.Pool 获取 TmallItemUpdateSimpleschemaGetAPIRequest
func GetTmallItemUpdateSimpleschemaGetAPIRequest() *TmallItemUpdateSimpleschemaGetAPIRequest {
	return poolTmallItemUpdateSimpleschemaGetAPIRequest.Get().(*TmallItemUpdateSimpleschemaGetAPIRequest)
}

// ReleaseTmallItemUpdateSimpleschemaGetAPIRequest 将 TmallItemUpdateSimpleschemaGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemUpdateSimpleschemaGetAPIRequest(v *TmallItemUpdateSimpleschemaGetAPIRequest) {
	v.Reset()
	poolTmallItemUpdateSimpleschemaGetAPIRequest.Put(v)
}
