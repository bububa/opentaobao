package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecyclePredeductSettleSyncAPIRequest 同步回收单线下打款明细 API请求
// taobao.recycle.prededuct.settle.sync
//
// 同步回收单线下打款明细
type TaobaoRecyclePredeductSettleSyncAPIRequest struct {
	model.Params
	// 线下打款金额，正数表示 服务商向用户支付的 金额，负数表示 服务商收取用户 的金额，单位为 分
	_offlineSettleFee int64
	// 回收单 ID
	_oldOrderId int64
	// 数据版本号，从1开始递增，天猫平台收到请求后会对这个字段做校验，只有比当前版本大的请求才会处理。
	_version int64
}

// NewTaobaoRecyclePredeductSettleSyncRequest 初始化TaobaoRecyclePredeductSettleSyncAPIRequest对象
func NewTaobaoRecyclePredeductSettleSyncRequest() *TaobaoRecyclePredeductSettleSyncAPIRequest {
	return &TaobaoRecyclePredeductSettleSyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRecyclePredeductSettleSyncAPIRequest) Reset() {
	r._offlineSettleFee = 0
	r._oldOrderId = 0
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRecyclePredeductSettleSyncAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.prededuct.settle.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRecyclePredeductSettleSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRecyclePredeductSettleSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineSettleFee is OfflineSettleFee Setter
// 线下打款金额，正数表示 服务商向用户支付的 金额，负数表示 服务商收取用户 的金额，单位为 分
func (r *TaobaoRecyclePredeductSettleSyncAPIRequest) SetOfflineSettleFee(_offlineSettleFee int64) error {
	r._offlineSettleFee = _offlineSettleFee
	r.Set("offline_settle_fee", _offlineSettleFee)
	return nil
}

// GetOfflineSettleFee OfflineSettleFee Getter
func (r TaobaoRecyclePredeductSettleSyncAPIRequest) GetOfflineSettleFee() int64 {
	return r._offlineSettleFee
}

// SetOldOrderId is OldOrderId Setter
// 回收单 ID
func (r *TaobaoRecyclePredeductSettleSyncAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaoRecyclePredeductSettleSyncAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}

// SetVersion is Version Setter
// 数据版本号，从1开始递增，天猫平台收到请求后会对这个字段做校验，只有比当前版本大的请求才会处理。
func (r *TaobaoRecyclePredeductSettleSyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TaobaoRecyclePredeductSettleSyncAPIRequest) GetVersion() int64 {
	return r._version
}

var poolTaobaoRecyclePredeductSettleSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRecyclePredeductSettleSyncRequest()
	},
}

// GetTaobaoRecyclePredeductSettleSyncRequest 从 sync.Pool 获取 TaobaoRecyclePredeductSettleSyncAPIRequest
func GetTaobaoRecyclePredeductSettleSyncAPIRequest() *TaobaoRecyclePredeductSettleSyncAPIRequest {
	return poolTaobaoRecyclePredeductSettleSyncAPIRequest.Get().(*TaobaoRecyclePredeductSettleSyncAPIRequest)
}

// ReleaseTaobaoRecyclePredeductSettleSyncAPIRequest 将 TaobaoRecyclePredeductSettleSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoRecyclePredeductSettleSyncAPIRequest(v *TaobaoRecyclePredeductSettleSyncAPIRequest) {
	v.Reset()
	poolTaobaoRecyclePredeductSettleSyncAPIRequest.Put(v)
}
