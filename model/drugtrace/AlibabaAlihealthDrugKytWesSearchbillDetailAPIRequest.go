package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwessearchbilldetailAPIRequest 查询单据详情 API请求
// alibaba.alihealth.drug.kyt.wes.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
type AlibabaalihealthdrugkytwessearchbilldetailAPIRequest struct {
	model.Params
	// 企业refEntId
	_refEntId string
	// licenseToken
	_licenseToken string
	// billCode
	_billCode string
	// 是否显示单据中的码( 1：显示    0：不显示 )
	_showCode string
}

// NewAlibabaalihealthdrugkytwessearchbilldetailRequest 初始化AlibabaalihealthdrugkytwessearchbilldetailAPIRequest对象
func NewAlibabaalihealthdrugkytwessearchbilldetailRequest() *AlibabaalihealthdrugkytwessearchbilldetailAPIRequest {
	return &AlibabaalihealthdrugkytwessearchbilldetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.searchbill.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业refEntId
func (r *AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// licenseToken
func (r *AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetBillCode is BillCode Setter
// billCode
func (r *AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetShowCode is ShowCode Setter
// 是否显示单据中的码( 1：显示    0：不显示 )
func (r *AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) SetShowCode(_showCode string) error {
	r._showCode = _showCode
	r.Set("show_code", _showCode)
	return nil
}

// GetShowCode ShowCode Getter
func (r AlibabaalihealthdrugkytwessearchbilldetailAPIRequest) GetShowCode() string {
	return r._showCode
}
