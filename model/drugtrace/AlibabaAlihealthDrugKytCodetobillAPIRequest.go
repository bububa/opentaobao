package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytCodetobillAPIRequest
通过追溯码查单据 API请求
alibaba.alihealth.drug.kyt.codetobill

通过追溯码查单据 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytCodetobillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.codetobill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytCodetobillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytCodetobillAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytCodetobillAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytCodetobillAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaAlihealthDrugKytCodetobillAPIRequest) GetCode() string {
	return r._code
}
