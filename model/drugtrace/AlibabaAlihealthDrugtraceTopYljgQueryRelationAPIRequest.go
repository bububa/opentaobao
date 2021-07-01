package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest
单码关联关系查询 API请求
alibaba.alihealth.drugtrace.top.yljg.query.relation

单码关联关系查询 */
type AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaAlihealthDrugtraceTopYljgQueryRelationRequest 初始化AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryRelationRequest() *AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest) GetCode() string {
	return r._code
}

// Set is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// Get DesRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}
