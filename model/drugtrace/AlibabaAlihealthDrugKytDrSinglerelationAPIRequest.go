package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrSinglerelationAPIRequest 多融单码关联关系查询 API请求
// alibaba.alihealth.drug.kyt.dr.singlerelation
//
// 单码关联关系查询
type AlibabaAlihealthDrugKytDrSinglerelationAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 追溯码
	_code string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaAlihealthDrugKytDrSinglerelationRequest 初始化AlibabaAlihealthDrugKytDrSinglerelationAPIRequest对象
func NewAlibabaAlihealthDrugKytDrSinglerelationRequest() *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest {
	return &AlibabaAlihealthDrugKytDrSinglerelationAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) Reset() {
	r._refEntId = ""
	r._code = ""
	r._desRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.singlerelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetCode() string {
	return r._code
}

// SetDesRefEntId is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// GetDesRefEntId DesRefEntId Getter
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}

var poolAlibabaAlihealthDrugKytDrSinglerelationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrSinglerelationRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrSinglerelationRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrSinglerelationAPIRequest
func GetAlibabaAlihealthDrugKytDrSinglerelationAPIRequest() *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest {
	return poolAlibabaAlihealthDrugKytDrSinglerelationAPIRequest.Get().(*AlibabaAlihealthDrugKytDrSinglerelationAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrSinglerelationAPIRequest 将 AlibabaAlihealthDrugKytDrSinglerelationAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrSinglerelationAPIRequest(v *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrSinglerelationAPIRequest.Put(v)
}
