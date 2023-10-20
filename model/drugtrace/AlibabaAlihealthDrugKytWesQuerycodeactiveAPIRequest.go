package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest 查询码是否激活 API请求
// alibaba.alihealth.drug.kyt.wes.querycodeactive
//
// 查询码是否激活
type AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 服务时间
	_licenseToken string
	// 码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes string
}

// NewAlibabaalihealthdrugkytwesquerycodeactiveRequest 初始化AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest对象
func NewAlibabaalihealthdrugkytwesquerycodeactiveRequest() *AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest {
	return &AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.querycodeactive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetCodes is Codes Setter
// 码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest) SetCodes(_codes string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest) GetCodes() string {
	return r._codes
}
