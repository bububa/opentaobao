package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqySinglerelationAPIRequest 单码关联关系查询 API请求
// alibaba.alihealth.drug.kyt.scqy.singlerelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
type AlibabaAlihealthDrugKytScqySinglerelationAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 追溯码
	_code string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaAlihealthDrugKytScqySinglerelationRequest 初始化AlibabaAlihealthDrugKytScqySinglerelationAPIRequest对象
func NewAlibabaAlihealthDrugKytScqySinglerelationRequest() *AlibabaAlihealthDrugKytScqySinglerelationAPIRequest {
	return &AlibabaAlihealthDrugKytScqySinglerelationAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) Reset() {
	r._refEntId = ""
	r._code = ""
	r._desRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.singlerelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) GetCode() string {
	return r._code
}

// SetDesRefEntId is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// GetDesRefEntId DesRefEntId Getter
func (r AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}

var poolAlibabaAlihealthDrugKytScqySinglerelationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytScqySinglerelationRequest()
	},
}

// GetAlibabaAlihealthDrugKytScqySinglerelationRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqySinglerelationAPIRequest
func GetAlibabaAlihealthDrugKytScqySinglerelationAPIRequest() *AlibabaAlihealthDrugKytScqySinglerelationAPIRequest {
	return poolAlibabaAlihealthDrugKytScqySinglerelationAPIRequest.Get().(*AlibabaAlihealthDrugKytScqySinglerelationAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytScqySinglerelationAPIRequest 将 AlibabaAlihealthDrugKytScqySinglerelationAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqySinglerelationAPIRequest(v *AlibabaAlihealthDrugKytScqySinglerelationAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqySinglerelationAPIRequest.Put(v)
}
