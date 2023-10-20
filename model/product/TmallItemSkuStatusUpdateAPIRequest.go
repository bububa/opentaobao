package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuStatusUpdateAPIRequest 商品sku状态更新 API请求
// tmall.item.sku.status.update
//
// 商品sku上下架状态更新
type TmallItemSkuStatusUpdateAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// sku状态信息
	_itemSkuStatus *ItemSkuStatus
}

// NewTmallItemSkuStatusUpdateRequest 初始化TmallItemSkuStatusUpdateAPIRequest对象
func NewTmallItemSkuStatusUpdateRequest() *TmallItemSkuStatusUpdateAPIRequest {
	return &TmallItemSkuStatusUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSkuStatusUpdateAPIRequest) Reset() {
	r._itemId = 0
	r._itemSkuStatus = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSkuStatusUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSkuStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSkuStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemSkuStatusUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSkuStatusUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemSkuStatus is ItemSkuStatus Setter
// sku状态信息
func (r *TmallItemSkuStatusUpdateAPIRequest) SetItemSkuStatus(_itemSkuStatus *ItemSkuStatus) error {
	r._itemSkuStatus = _itemSkuStatus
	r.Set("item_sku_status", _itemSkuStatus)
	return nil
}

// GetItemSkuStatus ItemSkuStatus Getter
func (r TmallItemSkuStatusUpdateAPIRequest) GetItemSkuStatus() *ItemSkuStatus {
	return r._itemSkuStatus
}

var poolTmallItemSkuStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSkuStatusUpdateRequest()
	},
}

// GetTmallItemSkuStatusUpdateRequest 从 sync.Pool 获取 TmallItemSkuStatusUpdateAPIRequest
func GetTmallItemSkuStatusUpdateAPIRequest() *TmallItemSkuStatusUpdateAPIRequest {
	return poolTmallItemSkuStatusUpdateAPIRequest.Get().(*TmallItemSkuStatusUpdateAPIRequest)
}

// ReleaseTmallItemSkuStatusUpdateAPIRequest 将 TmallItemSkuStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallItemSkuStatusUpdateAPIRequest(v *TmallItemSkuStatusUpdateAPIRequest) {
	v.Reset()
	poolTmallItemSkuStatusUpdateAPIRequest.Put(v)
}
