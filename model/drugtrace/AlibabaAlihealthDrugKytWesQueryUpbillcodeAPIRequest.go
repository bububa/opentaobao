package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest 通过一个码查询上游出库单 API请求
// alibaba.alihealth.drug.kyt.wes.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest struct {
	model.Params
	// 企业REF_ENT_ID （当前企业的唯一标识）
	_refEntId string
	// 服务时间
	_licenseToken string
	// 追溯码
	_code string
}

// NewAlibabaalihealthdrugkytwesqueryupbillcodeRequest 初始化AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest对象
func NewAlibabaalihealthdrugkytwesqueryupbillcodeRequest() *AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest {
	return &AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.query.upbillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业REF_ENT_ID （当前企业的唯一标识）
func (r *AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest) GetCode() string {
	return r._code
}
