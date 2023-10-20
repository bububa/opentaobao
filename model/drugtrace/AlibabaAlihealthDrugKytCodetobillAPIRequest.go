package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytCodetobillAPIRequest 通过追溯码查单据 API请求
// alibaba.alihealth.drug.kyt.codetobill
//
// 通过追溯码查单据
type AlibabaAlihealthDrugKytCodetobillAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 追溯码
	_code string
}

// NewAlibabaAlihealthDrugKytCodetobillRequest 初始化AlibabaAlihealthDrugKytCodetobillAPIRequest对象
func NewAlibabaAlihealthDrugKytCodetobillRequest() *AlibabaAlihealthDrugKytCodetobillAPIRequest {
	return &AlibabaAlihealthDrugKytCodetobillAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytCodetobillAPIRequest) Reset() {
	r._refEntId = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytCodetobillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.codetobill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytCodetobillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytCodetobillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytCodetobillAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytCodetobillAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytCodetobillAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytCodetobillAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaAlihealthDrugKytCodetobillAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytCodetobillRequest()
	},
}

// GetAlibabaAlihealthDrugKytCodetobillRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytCodetobillAPIRequest
func GetAlibabaAlihealthDrugKytCodetobillAPIRequest() *AlibabaAlihealthDrugKytCodetobillAPIRequest {
	return poolAlibabaAlihealthDrugKytCodetobillAPIRequest.Get().(*AlibabaAlihealthDrugKytCodetobillAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytCodetobillAPIRequest 将 AlibabaAlihealthDrugKytCodetobillAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytCodetobillAPIRequest(v *AlibabaAlihealthDrugKytCodetobillAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytCodetobillAPIRequest.Put(v)
}
