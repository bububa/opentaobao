package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest 检查输入的码之间是否有上下级关系 API请求
// alibaba.alihealth.drug.code.kyt.wes.checkcoderelation
//
// 检查输入的码之间是否有上下级关系
type AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest struct {
	model.Params
	// 企业refentid
	_refEntId string
	// 服务校验的token
	_licenseToken string
	// 多个码用英文逗号分隔
	_codes string
}

// NewAlibabaalihealthdrugcodekytwescheckcoderelationRequest 初始化AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest对象
func NewAlibabaalihealthdrugcodekytwescheckcoderelationRequest() *AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest {
	return &AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.wes.checkcoderelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业refentid
func (r *AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务校验的token
func (r *AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetCodes is Codes Setter
// 多个码用英文逗号分隔
func (r *AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest) SetCodes(_codes string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest) GetCodes() string {
	return r._codes
}
