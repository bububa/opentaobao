package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorecyclepredeductsettlesyncAPIRequest 同步回收单线下打款明细 API请求
// taobao.recycle.prededuct.settle.sync
//
// 同步回收单线下打款明细
type TaobaorecyclepredeductsettlesyncAPIRequest struct {
	model.Params
	// 线下打款金额，正数表示 服务商向用户支付的 金额，负数表示 服务商收取用户 的金额，单位为 分
	_offlineSettleFee int64
	// 回收单 ID
	_oldOrderId int64
	// 数据版本号，从1开始递增，天猫平台收到请求后会对这个字段做校验，只有比当前版本大的请求才会处理。
	_version int64
}

// NewTaobaorecyclepredeductsettlesyncRequest 初始化TaobaorecyclepredeductsettlesyncAPIRequest对象
func NewTaobaorecyclepredeductsettlesyncRequest() *TaobaorecyclepredeductsettlesyncAPIRequest {
	return &TaobaorecyclepredeductsettlesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorecyclepredeductsettlesyncAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.prededuct.settle.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorecyclepredeductsettlesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorecyclepredeductsettlesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineSettleFee is OfflineSettleFee Setter
// 线下打款金额，正数表示 服务商向用户支付的 金额，负数表示 服务商收取用户 的金额，单位为 分
func (r *TaobaorecyclepredeductsettlesyncAPIRequest) SetOfflineSettleFee(_offlineSettleFee int64) error {
	r._offlineSettleFee = _offlineSettleFee
	r.Set("offline_settle_fee", _offlineSettleFee)
	return nil
}

// GetOfflineSettleFee OfflineSettleFee Getter
func (r TaobaorecyclepredeductsettlesyncAPIRequest) GetOfflineSettleFee() int64 {
	return r._offlineSettleFee
}

// SetOldOrderId is OldOrderId Setter
// 回收单 ID
func (r *TaobaorecyclepredeductsettlesyncAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaorecyclepredeductsettlesyncAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}

// SetVersion is Version Setter
// 数据版本号，从1开始递增，天猫平台收到请求后会对这个字段做校验，只有比当前版本大的请求才会处理。
func (r *TaobaorecyclepredeductsettlesyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TaobaorecyclepredeductsettlesyncAPIRequest) GetVersion() int64 {
	return r._version
}
