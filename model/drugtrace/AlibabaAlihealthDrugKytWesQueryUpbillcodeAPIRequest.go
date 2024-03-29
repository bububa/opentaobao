package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest 通过一个码查询上游出库单 API请求
// alibaba.alihealth.drug.kyt.wes.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest struct {
	model.Params
	// 企业REF_ENT_ID （当前企业的唯一标识）
	_refEntId string
	// 服务时间
	_licenseToken string
	// 追溯码
	_code string
}

// NewAlibabaAlihealthDrugKytWesQueryUpbillcodeRequest 初始化AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest对象
func NewAlibabaAlihealthDrugKytWesQueryUpbillcodeRequest() *AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest {
	return &AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) Reset() {
	r._refEntId = ""
	r._licenseToken = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.query.upbillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业REF_ENT_ID （当前企业的唯一标识）
func (r *AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytWesQueryUpbillcodeRequest()
	},
}

// GetAlibabaAlihealthDrugKytWesQueryUpbillcodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest
func GetAlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest() *AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest {
	return poolAlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest.Get().(*AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest 将 AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest(v *AlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesQueryUpbillcodeAPIRequest.Put(v)
}
