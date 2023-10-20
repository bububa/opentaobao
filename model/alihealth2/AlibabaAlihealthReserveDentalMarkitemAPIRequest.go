package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthreservedentalmarkitemAPIRequest 标记商品是否可预约 API请求
// alibaba.alihealth.reserve.dental.markitem
//
// 标记商品是否可预约
type AlibabaalihealthreservedentalmarkitemAPIRequest struct {
	model.Params
	// 平台商品id
	_itemId int64
	// 是否可预约，1.可预约 0.不可预约
	_status int64
}

// NewAlibabaalihealthreservedentalmarkitemRequest 初始化AlibabaalihealthreservedentalmarkitemAPIRequest对象
func NewAlibabaalihealthreservedentalmarkitemRequest() *AlibabaalihealthreservedentalmarkitemAPIRequest {
	return &AlibabaalihealthreservedentalmarkitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthreservedentalmarkitemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.markitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthreservedentalmarkitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthreservedentalmarkitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 平台商品id
func (r *AlibabaalihealthreservedentalmarkitemAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaalihealthreservedentalmarkitemAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetStatus is Status Setter
// 是否可预约，1.可预约 0.不可预约
func (r *AlibabaalihealthreservedentalmarkitemAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaalihealthreservedentalmarkitemAPIRequest) GetStatus() int64 {
	return r._status
}
