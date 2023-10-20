package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecycleOrderFulfillSyncAPIRequest 同步回收单最终履约方式 API请求
// taobao.recycle.order.fulfill.sync
//
// 同步回收单最终履约方式
type TaobaoRecycleOrderFulfillSyncAPIRequest struct {
	model.Params
	// 取送一体时这个字段传 1，否则传 0
	_fulfillType int64
	// 回收单订单 ID
	_oldOrderId int64
}

// NewTaobaoRecycleOrderFulfillSyncRequest 初始化TaobaoRecycleOrderFulfillSyncAPIRequest对象
func NewTaobaoRecycleOrderFulfillSyncRequest() *TaobaoRecycleOrderFulfillSyncAPIRequest {
	return &TaobaoRecycleOrderFulfillSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRecycleOrderFulfillSyncAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.order.fulfill.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRecycleOrderFulfillSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRecycleOrderFulfillSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillType is FulfillType Setter
// 取送一体时这个字段传 1，否则传 0
func (r *TaobaoRecycleOrderFulfillSyncAPIRequest) SetFulfillType(_fulfillType int64) error {
	r._fulfillType = _fulfillType
	r.Set("fulfill_type", _fulfillType)
	return nil
}

// GetFulfillType FulfillType Getter
func (r TaobaoRecycleOrderFulfillSyncAPIRequest) GetFulfillType() int64 {
	return r._fulfillType
}

// SetOldOrderId is OldOrderId Setter
// 回收单订单 ID
func (r *TaobaoRecycleOrderFulfillSyncAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaoRecycleOrderFulfillSyncAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}
