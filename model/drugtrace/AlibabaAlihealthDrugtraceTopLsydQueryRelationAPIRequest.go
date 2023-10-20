package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest 单码关联关系查询 API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.relation
//
// 单码关联关系查询
type AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 追溯码
	_code string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaAlihealthDrugtraceTopLsydQueryRelationRequest 初始化AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryRelationRequest() *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) Reset() {
	r._refEntId = ""
	r._code = ""
	r._desRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) GetCode() string {
	return r._code
}

// SetDesRefEntId is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// GetDesRefEntId DesRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopLsydQueryRelationRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryRelationRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest
func GetAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest() *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest 将 AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest(v *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest.Put(v)
}
