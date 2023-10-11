package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest 检查输入的码之间是否有上下级关系 API请求
// alibaba.alihealth.drug.code.kyt.wes.checkcoderelation
//
// 检查输入的码之间是否有上下级关系
type AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest struct {
	model.Params
	// 企业refentid
	_refEntId string
	// 服务校验的token
	_licenseToken string
	// 多个码用英文逗号分隔
	_codes string
}

// NewAlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest 初始化AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest() *AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest {
	return &AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.wes.checkcoderelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业refentid
func (r *AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务校验的token
func (r *AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetCodes is Codes Setter
// 多个码用英文逗号分隔
func (r *AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest) SetCodes(_codes string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest) GetCodes() string {
	return r._codes
}
