package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest 单码关联关系查询 API请求
// alibaba.alihealth.drugtrace.top.yljg.query.relation
//
// 单码关联关系查询
type AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 追溯码
	_code string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaalihealthdrugtracetopyljgqueryrelationRequest 初始化AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest对象
func NewAlibabaalihealthdrugtracetopyljgqueryrelationRequest() *AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest {
	return &AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest) GetCode() string {
	return r._code
}

// SetDesRefEntId is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// GetDesRefEntId DesRefEntId Getter
func (r AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}
