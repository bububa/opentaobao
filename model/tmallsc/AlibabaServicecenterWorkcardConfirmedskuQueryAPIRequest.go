package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest 查询确认履行的服务项 API请求
// alibaba.servicecenter.workcard.confirmedsku.query
//
// 查询确认履行的服务项
type AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest struct {
	model.Params
	// 真实履约服务商Nick（非ERP系统不要填写）
	_realTpNick string
	// 工单id
	_workcardId int64
}

// NewAlibabaServicecenterWorkcardConfirmedskuQueryRequest 初始化AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest对象
func NewAlibabaServicecenterWorkcardConfirmedskuQueryRequest() *AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest {
	return &AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.confirmedsku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRealTpNick is RealTpNick Setter
// 真实履约服务商Nick（非ERP系统不要填写）
func (r *AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
