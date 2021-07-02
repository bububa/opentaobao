package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardCancelAPIRequest 服务平台工单取消接口 API请求
// alibaba.servicecenter.workcard.cancel
//
// 取消服务工单
type AlibabaServicecenterWorkcardCancelAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 取消备注
	_memo string
	// 服务单id
	_serviceOrderId int64
	// 真实服务商昵称
	_realTpNick string
}

// NewAlibabaServicecenterWorkcardCancelRequest 初始化AlibabaServicecenterWorkcardCancelAPIRequest对象
func NewAlibabaServicecenterWorkcardCancelRequest() *AlibabaServicecenterWorkcardCancelAPIRequest {
	return &AlibabaServicecenterWorkcardCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterWorkcardCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterWorkcardCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WorkcardId Setter
// 工单id
func (r *AlibabaServicecenterWorkcardCancelAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// Get WorkcardId Getter
func (r AlibabaServicecenterWorkcardCancelAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// Set is Memo Setter
// 取消备注
func (r *AlibabaServicecenterWorkcardCancelAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// Get Memo Getter
func (r AlibabaServicecenterWorkcardCancelAPIRequest) GetMemo() string {
	return r._memo
}

// Set is ServiceOrderId Setter
// 服务单id
func (r *AlibabaServicecenterWorkcardCancelAPIRequest) SetServiceOrderId(_serviceOrderId int64) error {
	r._serviceOrderId = _serviceOrderId
	r.Set("service_order_id", _serviceOrderId)
	return nil
}

// Get ServiceOrderId Getter
func (r AlibabaServicecenterWorkcardCancelAPIRequest) GetServiceOrderId() int64 {
	return r._serviceOrderId
}

// Set is RealTpNick Setter
// 真实服务商昵称
func (r *AlibabaServicecenterWorkcardCancelAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// Get RealTpNick Getter
func (r AlibabaServicecenterWorkcardCancelAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}
