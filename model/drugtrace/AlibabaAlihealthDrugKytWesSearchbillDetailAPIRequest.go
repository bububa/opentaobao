package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest 查询单据详情 API请求
// alibaba.alihealth.drug.kyt.wes.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
type AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest struct {
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

// NewAlibabaAlihealthDrugKytWesSearchbillDetailRequest 初始化AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest对象
func NewAlibabaAlihealthDrugKytWesSearchbillDetailRequest() *AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest {
	return &AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) Reset() {
	r._refEntId = ""
	r._licenseToken = ""
	r._billCode = ""
	r._showCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.searchbill.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// licenseToken
func (r *AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetBillCode is BillCode Setter
// billCode
func (r *AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetShowCode is ShowCode Setter
// 是否显示单据中的码( 1：显示    0：不显示 )
func (r *AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) SetShowCode(_showCode string) error {
	r._showCode = _showCode
	r.Set("show_code", _showCode)
	return nil
}

// GetShowCode ShowCode Getter
func (r AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) GetShowCode() string {
	return r._showCode
}

var poolAlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytWesSearchbillDetailRequest()
	},
}

// GetAlibabaAlihealthDrugKytWesSearchbillDetailRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest
func GetAlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest() *AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest {
	return poolAlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest.Get().(*AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest 将 AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest(v *AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest.Put(v)
}
