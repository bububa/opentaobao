package xhotelofficial

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfficialSettleCancelAPIRequest 官网信用住取消结账 API请求
// taobao.xhotel.order.official.settle.cancel
//
// 用于官网信用住取消结账
type TaobaoXhotelOrderOfficialSettleCancelAPIRequest struct {
	model.Params
	// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
	_tid int64
	// 取消结账的原因
	_reason string
	// 外部订单号，和tid二选一必填
	_outId string
	// 暂时无意义，无需传入
	_notifyUrl string
	// 请求流水号
	_outUuid string
}

// NewTaobaoXhotelOrderOfficialSettleCancelRequest 初始化TaobaoXhotelOrderOfficialSettleCancelAPIRequest对象
func NewTaobaoXhotelOrderOfficialSettleCancelRequest() *TaobaoXhotelOrderOfficialSettleCancelAPIRequest {
	return &TaobaoXhotelOrderOfficialSettleCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.official.settle.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Tid Setter
// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetTid() int64 {
	return r._tid
}

// Set is Reason Setter
// 取消结账的原因
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// Get Reason Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetReason() string {
	return r._reason
}

// Set is OutId Setter
// 外部订单号，和tid二选一必填
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// Get OutId Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetOutId() string {
	return r._outId
}

// Set is NotifyUrl Setter
// 暂时无意义，无需传入
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetNotifyUrl(_notifyUrl string) error {
	r._notifyUrl = _notifyUrl
	r.Set("notify_url", _notifyUrl)
	return nil
}

// Get NotifyUrl Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetNotifyUrl() string {
	return r._notifyUrl
}

// Set is OutUuid Setter
// 请求流水号
func (r *TaobaoXhotelOrderOfficialSettleCancelAPIRequest) SetOutUuid(_outUuid string) error {
	r._outUuid = _outUuid
	r.Set("out_uuid", _outUuid)
	return nil
}

// Get OutUuid Getter
func (r TaobaoXhotelOrderOfficialSettleCancelAPIRequest) GetOutUuid() string {
	return r._outUuid
}
