package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterworkcardcancelAPIRequest 服务平台工单取消接口 API请求
// alibaba.servicecenter.workcard.cancel
//
// 取消服务工单
type AlibabaservicecenterworkcardcancelAPIRequest struct {
	model.Params
	// 取消备注
	_memo string
	// 真实服务商昵称
	_realTpNick string
	// 原因desc
	_reasonDesc string
	// 工单id
	_workcardId int64
	// 服务单id
	_serviceOrderId int64
	// 请求来源类型
	_type int64
	// 原因code
	_reasonCode int64
}

// NewAlibabaservicecenterworkcardcancelRequest 初始化AlibabaservicecenterworkcardcancelAPIRequest对象
func NewAlibabaservicecenterworkcardcancelRequest() *AlibabaservicecenterworkcardcancelAPIRequest {
	return &AlibabaservicecenterworkcardcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 取消备注
func (r *AlibabaservicecenterworkcardcancelAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetMemo() string {
	return r._memo
}

// SetRealTpNick is RealTpNick Setter
// 真实服务商昵称
func (r *AlibabaservicecenterworkcardcancelAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetReasonDesc is ReasonDesc Setter
// 原因desc
func (r *AlibabaservicecenterworkcardcancelAPIRequest) SetReasonDesc(_reasonDesc string) error {
	r._reasonDesc = _reasonDesc
	r.Set("reason_desc", _reasonDesc)
	return nil
}

// GetReasonDesc ReasonDesc Getter
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetReasonDesc() string {
	return r._reasonDesc
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *AlibabaservicecenterworkcardcancelAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetServiceOrderId is ServiceOrderId Setter
// 服务单id
func (r *AlibabaservicecenterworkcardcancelAPIRequest) SetServiceOrderId(_serviceOrderId int64) error {
	r._serviceOrderId = _serviceOrderId
	r.Set("service_order_id", _serviceOrderId)
	return nil
}

// GetServiceOrderId ServiceOrderId Getter
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetServiceOrderId() int64 {
	return r._serviceOrderId
}

// SetType is Type Setter
// 请求来源类型
func (r *AlibabaservicecenterworkcardcancelAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetType() int64 {
	return r._type
}

// SetReasonCode is ReasonCode Setter
// 原因code
func (r *AlibabaservicecenterworkcardcancelAPIRequest) SetReasonCode(_reasonCode int64) error {
	r._reasonCode = _reasonCode
	r.Set("reason_code", _reasonCode)
	return nil
}

// GetReasonCode ReasonCode Getter
func (r AlibabaservicecenterworkcardcancelAPIRequest) GetReasonCode() int64 {
	return r._reasonCode
}
