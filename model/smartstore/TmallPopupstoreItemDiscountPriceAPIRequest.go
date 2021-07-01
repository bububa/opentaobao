package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPopupstoreItemDiscountPriceAPIRequest
商品优惠价格查询 API请求
tmall.popupstore.item.discount.price

商品优惠价格查询 */
type TmallPopupstoreItemDiscountPriceAPIRequest struct {
	model.Params
	// 商品id列表
	_itemIds []int64
}

// NewTmallPopupstoreItemDiscountPriceRequest 初始化TmallPopupstoreItemDiscountPriceAPIRequest对象
func NewTmallPopupstoreItemDiscountPriceRequest() *TmallPopupstoreItemDiscountPriceAPIRequest {
	return &TmallPopupstoreItemDiscountPriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPopupstoreItemDiscountPriceAPIRequest) GetApiMethodName() string {
	return "tmall.popupstore.item.discount.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPopupstoreItemDiscountPriceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemIds Setter
// 商品id列表
func (r *TmallPopupstoreItemDiscountPriceAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// Get ItemIds Getter
func (r TmallPopupstoreItemDiscountPriceAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}
