package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfflineSettleCancelAPIRequest 线下信用住取消结账专用接口 API请求
// taobao.xhotel.order.offline.settle.cancel
//
// 线下信用住取消结账专用接口
type TaobaoXhotelOrderOfflineSettleCancelAPIRequest struct {
	model.Params
	// 取消结账的原因
	_reason string
	// 外部订单号，和tid二选一必填（建议都写入）
	_outId string
	// 暂时无意义，无需传入
	_notifyUrl string
	// 请求流水号
	_outUuid string
	// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
	_tid int64
}

// NewTaobaoXhotelOrderOfflineSettleCancelRequest 初始化TaobaoXhotelOrderOfflineSettleCancelAPIRequest对象
func NewTaobaoXhotelOrderOfflineSettleCancelRequest() *TaobaoXhotelOrderOfflineSettleCancelAPIRequest {
	return &TaobaoXhotelOrderOfflineSettleCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfflineSettleCancelAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.offline.settle.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfflineSettleCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelOrderOfflineSettleCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReason is Reason Setter
// 取消结账的原因
func (r *TaobaoXhotelOrderOfflineSettleCancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoXhotelOrderOfflineSettleCancelAPIRequest) GetReason() string {
	return r._reason
}

// SetOutId is OutId Setter
// 外部订单号，和tid二选一必填（建议都写入）
func (r *TaobaoXhotelOrderOfflineSettleCancelAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoXhotelOrderOfflineSettleCancelAPIRequest) GetOutId() string {
	return r._outId
}

// SetNotifyUrl is NotifyUrl Setter
// 暂时无意义，无需传入
func (r *TaobaoXhotelOrderOfflineSettleCancelAPIRequest) SetNotifyUrl(_notifyUrl string) error {
	r._notifyUrl = _notifyUrl
	r.Set("notify_url", _notifyUrl)
	return nil
}

// GetNotifyUrl NotifyUrl Getter
func (r TaobaoXhotelOrderOfflineSettleCancelAPIRequest) GetNotifyUrl() string {
	return r._notifyUrl
}

// SetOutUuid is OutUuid Setter
// 请求流水号
func (r *TaobaoXhotelOrderOfflineSettleCancelAPIRequest) SetOutUuid(_outUuid string) error {
	r._outUuid = _outUuid
	r.Set("out_uuid", _outUuid)
	return nil
}

// GetOutUuid OutUuid Getter
func (r TaobaoXhotelOrderOfflineSettleCancelAPIRequest) GetOutUuid() string {
	return r._outUuid
}

// SetTid is Tid Setter
// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
func (r *TaobaoXhotelOrderOfflineSettleCancelAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelOrderOfflineSettleCancelAPIRequest) GetTid() int64 {
	return r._tid
}
