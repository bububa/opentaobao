package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemskustatusupdateAPIRequest 商品sku状态更新 API请求
// tmall.item.sku.status.update
//
// 商品sku上下架状态更新
type TmallitemskustatusupdateAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// sku状态信息
	_itemSkuStatus *ItemSkuStatus
}

// NewTmallitemskustatusupdateRequest 初始化TmallitemskustatusupdateAPIRequest对象
func NewTmallitemskustatusupdateRequest() *TmallitemskustatusupdateAPIRequest {
	return &TmallitemskustatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemskustatusupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemskustatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemskustatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallitemskustatusupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemskustatusupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemSkuStatus is ItemSkuStatus Setter
// sku状态信息
func (r *TmallitemskustatusupdateAPIRequest) SetItemSkuStatus(_itemSkuStatus *ItemSkuStatus) error {
	r._itemSkuStatus = _itemSkuStatus
	r.Set("item_sku_status", _itemSkuStatus)
	return nil
}

// GetItemSkuStatus ItemSkuStatus Getter
func (r TmallitemskustatusupdateAPIRequest) GetItemSkuStatus() *ItemSkuStatus {
	return r._itemSkuStatus
}
