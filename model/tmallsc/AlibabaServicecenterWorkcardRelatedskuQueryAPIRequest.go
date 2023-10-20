package tmallsc

import (
	"net/url"
	"sync"

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
	// 是否根据类目过滤维修项
	_isFilterByCategory bool
}

// NewAlibabaServicecenterWorkcardRelatedskuQueryRequest 初始化AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest对象
func NewAlibabaServicecenterWorkcardRelatedskuQueryRequest() *AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest {
	return &AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) Reset() {
	r._realTpNick = ""
	r._workcardId = 0
	r._isFilterByCategory = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.relatedsku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

// SetIsFilterByCategory is IsFilterByCategory Setter
// 是否根据类目过滤维修项
func (r *AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) SetIsFilterByCategory(_isFilterByCategory bool) error {
	r._isFilterByCategory = _isFilterByCategory
	r.Set("is_filter_by_category", _isFilterByCategory)
	return nil
}

// GetIsFilterByCategory IsFilterByCategory Getter
func (r AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) GetIsFilterByCategory() bool {
	return r._isFilterByCategory
}

var poolAlibabaServicecenterWorkcardRelatedskuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaServicecenterWorkcardRelatedskuQueryRequest()
	},
}

// GetAlibabaServicecenterWorkcardRelatedskuQueryRequest 从 sync.Pool 获取 AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest
func GetAlibabaServicecenterWorkcardRelatedskuQueryAPIRequest() *AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest {
	return poolAlibabaServicecenterWorkcardRelatedskuQueryAPIRequest.Get().(*AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest)
}

// ReleaseAlibabaServicecenterWorkcardRelatedskuQueryAPIRequest 将 AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaServicecenterWorkcardRelatedskuQueryAPIRequest(v *AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest) {
	v.Reset()
	poolAlibabaServicecenterWorkcardRelatedskuQueryAPIRequest.Put(v)
}
