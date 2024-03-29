package xhotelofficial

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfficialSettleCancelAPIRequest 官网信用住取消结账 API请求
// taobao.xhotel.order.official.settle.cancel
//
// 用于官网信用住取消结账
type TaobaoXhotelOrderOfficialSettleCancelAPIRequest struct {
	model.Params
	// 取消结账的原因
	_reason string
	// 外部订单号，和tid二选一必填
	_outId string
	// 暂时无意义，无需传入
	_notifyUrl string
	// 请求流水号
	_outUuid string
	// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
	_tid int64
}

// NewTaobaoXhotelOrderOfficialSettleCancelRequest 初始化TaobaoXhotelOrderOfficialSettleCancelAPIRequest对象
func NewTaobaoXhotelOrderOfficialSettleCancelRequest() *TaobaoXhotelOrderOfficialSettleCancelAPIRequest {
	return &TaobaoXhotelOrderOfficialSettleCancelAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) Reset() {
	r._reason = ""
	r._outId = ""
	r._notifyUrl = ""
	r._outUuid = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.official.settle.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReason is Reason Setter
// 取消结账的原因
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetReason() string {
	return r._reason
}

// SetOutId is OutId Setter
// 外部订单号，和tid二选一必填
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetOutId() string {
	return r._outId
}

// SetNotifyUrl is NotifyUrl Setter
// 暂时无意义，无需传入
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetNotifyUrl(_notifyUrl string) error {
	r._notifyUrl = _notifyUrl
	r.Set("notify_url", _notifyUrl)
	return nil
}

// GetNotifyUrl NotifyUrl Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetNotifyUrl() string {
	return r._notifyUrl
}

// SetOutUuid is OutUuid Setter
// 请求流水号
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetOutUuid(_outUuid string) error {
	r._outUuid = _outUuid
	r.Set("out_uuid", _outUuid)
	return nil
}

// GetOutUuid OutUuid Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetOutUuid() string {
	return r._outUuid
}

// SetTid is Tid Setter
// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoXhotelOrderOfficialSettleCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelOrderOfficialSettleCancelRequest()
	},
}

// GetTaobaoXhotelOrderOfficialSettleCancelRequest 从 sync.Pool 获取 TaobaoXhotelOrderOfficialSettleCancelAPIRequest
func GetTaobaoXhotelOrderOfficialSettleCancelAPIRequest() *TaobaoXhotelOrderOfficialSettleCancelAPIRequest {
	return poolTaobaoXhotelOrderOfficialSettleCancelAPIRequest.Get().(*TaobaoXhotelOrderOfficialSettleCancelAPIRequest)
}

// ReleaseTaobaoXhotelOrderOfficialSettleCancelAPIRequest 将 TaobaoXhotelOrderOfficialSettleCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelOrderOfficialSettleCancelAPIRequest(v *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) {
	v.Reset()
	poolTaobaoXhotelOrderOfficialSettleCancelAPIRequest.Put(v)
}
