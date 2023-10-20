package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterworkcardrelatedskuqueryAPIRequest 查询工单关联的服务项 API请求
// alibaba.servicecenter.workcard.relatedsku.query
//
// 查询工单关联的服务项
type AlibabaservicecenterworkcardrelatedskuqueryAPIRequest struct {
	model.Params
	// 真实履约服务商nick(非erp系统不要填写)
	_realTpNick string
	// 工单id
	_workcardId int64
	// 是否根据类目过滤维修项
	_isFilterByCategory bool
}

// NewAlibabaservicecenterworkcardrelatedskuqueryRequest 初始化AlibabaservicecenterworkcardrelatedskuqueryAPIRequest对象
func NewAlibabaservicecenterworkcardrelatedskuqueryRequest() *AlibabaservicecenterworkcardrelatedskuqueryAPIRequest {
	return &AlibabaservicecenterworkcardrelatedskuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterworkcardrelatedskuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.relatedsku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterworkcardrelatedskuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterworkcardrelatedskuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRealTpNick is RealTpNick Setter
// 真实履约服务商nick(非erp系统不要填写)
func (r *AlibabaservicecenterworkcardrelatedskuqueryAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r AlibabaservicecenterworkcardrelatedskuqueryAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *AlibabaservicecenterworkcardrelatedskuqueryAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaservicecenterworkcardrelatedskuqueryAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetIsFilterByCategory is IsFilterByCategory Setter
// 是否根据类目过滤维修项
func (r *AlibabaservicecenterworkcardrelatedskuqueryAPIRequest) SetIsFilterByCategory(_isFilterByCategory bool) error {
	r._isFilterByCategory = _isFilterByCategory
	r.Set("is_filter_by_category", _isFilterByCategory)
	return nil
}

// GetIsFilterByCategory IsFilterByCategory Getter
func (r AlibabaservicecenterworkcardrelatedskuqueryAPIRequest) GetIsFilterByCategory() bool {
	return r._isFilterByCategory
}
