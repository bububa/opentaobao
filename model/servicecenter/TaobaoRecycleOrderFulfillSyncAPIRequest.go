package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorecycleorderfulfillsyncAPIRequest 同步回收单最终履约方式 API请求
// taobao.recycle.order.fulfill.sync
//
// 同步回收单最终履约方式
type TaobaorecycleorderfulfillsyncAPIRequest struct {
	model.Params
	// 取送一体时这个字段传 1，否则传 0
	_fulfillType int64
	// 回收单订单 ID
	_oldOrderId int64
}

// NewTaobaorecycleorderfulfillsyncRequest 初始化TaobaorecycleorderfulfillsyncAPIRequest对象
func NewTaobaorecycleorderfulfillsyncRequest() *TaobaorecycleorderfulfillsyncAPIRequest {
	return &TaobaorecycleorderfulfillsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorecycleorderfulfillsyncAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.order.fulfill.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorecycleorderfulfillsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorecycleorderfulfillsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillType is FulfillType Setter
// 取送一体时这个字段传 1，否则传 0
func (r *TaobaorecycleorderfulfillsyncAPIRequest) SetFulfillType(_fulfillType int64) error {
	r._fulfillType = _fulfillType
	r.Set("fulfill_type", _fulfillType)
	return nil
}

// GetFulfillType FulfillType Getter
func (r TaobaorecycleorderfulfillsyncAPIRequest) GetFulfillType() int64 {
	return r._fulfillType
}

// SetOldOrderId is OldOrderId Setter
// 回收单订单 ID
func (r *TaobaorecycleorderfulfillsyncAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaorecycleorderfulfillsyncAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}
