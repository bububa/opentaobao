package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest wes码关联关系查询 API请求
// alibaba.alihealth.drug.code.kyt.wes.querycoderelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
type AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// license_token
	_licenseToken string
	// 追溯码
	_code string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest 初始化AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest() *AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest {
	return &AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.wes.querycoderelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// license_token
func (r *AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) GetCode() string {
	return r._code
}

// SetDesRefEntId is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// GetDesRefEntId DesRefEntId Getter
func (r AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}
