package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalMarkitemAPIRequest 标记商品是否可预约 API请求
// alibaba.alihealth.reserve.dental.markitem
//
// 标记商品是否可预约
type AlibabaAlihealthReserveDentalMarkitemAPIRequest struct {
	model.Params
	// 平台商品id
	_itemId int64
	// 是否可预约，1.可预约 0.不可预约
	_status int64
}

// NewAlibabaAlihealthReserveDentalMarkitemRequest 初始化AlibabaAlihealthReserveDentalMarkitemAPIRequest对象
func NewAlibabaAlihealthReserveDentalMarkitemRequest() *AlibabaAlihealthReserveDentalMarkitemAPIRequest {
	return &AlibabaAlihealthReserveDentalMarkitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalMarkitemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.markitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalMarkitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 平台商品id
func (r *AlibabaAlihealthReserveDentalMarkitemAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlibabaAlihealthReserveDentalMarkitemAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is Status Setter
// 是否可预约，1.可预约 0.不可预约
func (r *AlibabaAlihealthReserveDentalMarkitemAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaAlihealthReserveDentalMarkitemAPIRequest) GetStatus() int64 {
	return r._status
}
