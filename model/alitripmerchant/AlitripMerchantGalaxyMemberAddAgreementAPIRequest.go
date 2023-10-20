package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymemberaddagreementAPIRequest 添加用户协议记录接口 API请求
// alitrip.merchant.galaxy.member.add.agreement
//
// 记录用户是否授权协议
type AlitripmerchantgalaxymemberaddagreementAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 微信code(实时获取)
	_code string
	// 隐私协议枚举type
	_privacyAgreement bool
	// 会员协议枚举
	_dataExportAgreement bool
}

// NewAlitripmerchantgalaxymemberaddagreementRequest 初始化AlitripmerchantgalaxymemberaddagreementAPIRequest对象
func NewAlitripmerchantgalaxymemberaddagreementRequest() *AlitripmerchantgalaxymemberaddagreementAPIRequest {
	return &AlitripmerchantgalaxymemberaddagreementAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymemberaddagreementAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.add.agreement"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymemberaddagreementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymemberaddagreementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxymemberaddagreementAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymemberaddagreementAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetCode is Code Setter
// 微信code(实时获取)
func (r *AlitripmerchantgalaxymemberaddagreementAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripmerchantgalaxymemberaddagreementAPIRequest) GetCode() string {
	return r._code
}

// SetPrivacyAgreement is PrivacyAgreement Setter
// 隐私协议枚举type
func (r *AlitripmerchantgalaxymemberaddagreementAPIRequest) SetPrivacyAgreement(_privacyAgreement bool) error {
	r._privacyAgreement = _privacyAgreement
	r.Set("privacy_agreement", _privacyAgreement)
	return nil
}

// GetPrivacyAgreement PrivacyAgreement Getter
func (r AlitripmerchantgalaxymemberaddagreementAPIRequest) GetPrivacyAgreement() bool {
	return r._privacyAgreement
}

// SetDataExportAgreement is DataExportAgreement Setter
// 会员协议枚举
func (r *AlitripmerchantgalaxymemberaddagreementAPIRequest) SetDataExportAgreement(_dataExportAgreement bool) error {
	r._dataExportAgreement = _dataExportAgreement
	r.Set("data_export_agreement", _dataExportAgreement)
	return nil
}

// GetDataExportAgreement DataExportAgreement Getter
func (r AlitripmerchantgalaxymemberaddagreementAPIRequest) GetDataExportAgreement() bool {
	return r._dataExportAgreement
}
