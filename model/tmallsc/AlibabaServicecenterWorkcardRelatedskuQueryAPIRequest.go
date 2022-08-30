package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest 查询工单关联的服务项 API请求
// alibaba.servicecenter.workcard.relatedsku.query
//
// 查询工单关联的服务项
type AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest struct {
	model.Params
	// 真实履约服务商nick(非erp系统不要填写)
	_realTpNick string
	// 工单id
	_workcardId int64
}

// NewAlibabaServicecenterWorkcardRelatedskuQueryRequest 初始化AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest对象
func NewAlibabaServicecenterWorkcardRelatedskuQueryRequest() *AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest {
	return &AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.relatedsku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRealTpNick is RealTpNick Setter
// 真实履约服务商nick(非erp系统不要填写)
func (r *AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
