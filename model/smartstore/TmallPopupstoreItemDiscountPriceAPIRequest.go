package smartstore

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPopupstoreItemDiscountPriceAPIRequest 商品优惠价格查询 API请求
// tmall.popupstore.item.discount.price
//
// 商品优惠价格查询
type TmallPopupstoreItemDiscountPriceAPIRequest struct {
	model.Params
	// 商品id列表
	_itemIds []string
}

// NewTmallPopupstoreItemDiscountPriceRequest 初始化TmallPopupstoreItemDiscountPriceAPIRequest对象
func NewTmallPopupstoreItemDiscountPriceRequest() *TmallPopupstoreItemDiscountPriceAPIRequest {
	return &TmallPopupstoreItemDiscountPriceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallPopupstoreItemDiscountPriceAPIRequest) Reset() {
	r._itemIds = r._itemIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPopupstoreItemDiscountPriceAPIRequest) GetApiMethodName() string {
	return "tmall.popupstore.item.discount.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPopupstoreItemDiscountPriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallPopupstoreItemDiscountPriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 商品id列表
func (r *TmallPopupstoreItemDiscountPriceAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TmallPopupstoreItemDiscountPriceAPIRequest) GetItemIds() []string {
	return r._itemIds
}

var poolTmallPopupstoreItemDiscountPriceAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallPopupstoreItemDiscountPriceRequest()
	},
}

// GetTmallPopupstoreItemDiscountPriceRequest 从 sync.Pool 获取 TmallPopupstoreItemDiscountPriceAPIRequest
func GetTmallPopupstoreItemDiscountPriceAPIRequest() *TmallPopupstoreItemDiscountPriceAPIRequest {
	return poolTmallPopupstoreItemDiscountPriceAPIRequest.Get().(*TmallPopupstoreItemDiscountPriceAPIRequest)
}

// ReleaseTmallPopupstoreItemDiscountPriceAPIRequest 将 TmallPopupstoreItemDiscountPriceAPIRequest 放入 sync.Pool
func ReleaseTmallPopupstoreItemDiscountPriceAPIRequest(v *TmallPopupstoreItemDiscountPriceAPIRequest) {
	v.Reset()
	poolTmallPopupstoreItemDiscountPriceAPIRequest.Put(v)
}
