package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSinglerelationAPIRequest 单码关联关系查询，通过一个码查询这个码下的所有子码 API请求
// alibaba.alihealth.drug.kyt.singlerelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
type AlibabaAlihealthDrugKytSinglerelationAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 追溯码
	_code string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaAlihealthDrugKytSinglerelationRequest 初始化AlibabaAlihealthDrugKytSinglerelationAPIRequest对象
func NewAlibabaAlihealthDrugKytSinglerelationRequest() *AlibabaAlihealthDrugKytSinglerelationAPIRequest {
	return &AlibabaAlihealthDrugKytSinglerelationAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytSinglerelationAPIRequest) Reset() {
	r._refEntId = ""
	r._code = ""
	r._desRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.singlerelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytSinglerelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytSinglerelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetCode() string {
	return r._code
}

// SetDesRefEntId is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugKytSinglerelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// GetDesRefEntId DesRefEntId Getter
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}

var poolAlibabaAlihealthDrugKytSinglerelationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytSinglerelationRequest()
	},
}

// GetAlibabaAlihealthDrugKytSinglerelationRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytSinglerelationAPIRequest
func GetAlibabaAlihealthDrugKytSinglerelationAPIRequest() *AlibabaAlihealthDrugKytSinglerelationAPIRequest {
	return poolAlibabaAlihealthDrugKytSinglerelationAPIRequest.Get().(*AlibabaAlihealthDrugKytSinglerelationAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytSinglerelationAPIRequest 将 AlibabaAlihealthDrugKytSinglerelationAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSinglerelationAPIRequest(v *AlibabaAlihealthDrugKytSinglerelationAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSinglerelationAPIRequest.Put(v)
}
