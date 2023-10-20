package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest 单码关联关系查询 API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.relation
//
// 单码关联关系查询
type AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 追溯码
	_code string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaalihealthdrugtracetoplsydqueryrelationRequest 初始化AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest对象
func NewAlibabaalihealthdrugtracetoplsydqueryrelationRequest() *AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest {
	return &AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest) GetCode() string {
	return r._code
}

// SetDesRefEntId is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// GetDesRefEntId DesRefEntId Getter
func (r AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}
