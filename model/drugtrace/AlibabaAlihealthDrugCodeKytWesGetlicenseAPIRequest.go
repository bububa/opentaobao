package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest 获取licenseToken API请求
// alibaba.alihealth.drug.code.kyt.wes.getlicense
//
// 获取licenseToken
type AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// license
	_license string
}

// NewAlibabaAlihealthDrugCodeKytWesGetlicenseRequest 初始化AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytWesGetlicenseRequest() *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest {
	return &AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest) Reset() {
	r._refEntId = ""
	r._license = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.wes.getlicense"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicense is License Setter
// license
func (r *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest) SetLicense(_license string) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest) GetLicense() string {
	return r._license
}

var poolAlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeKytWesGetlicenseRequest()
	},
}

// GetAlibabaAlihealthDrugCodeKytWesGetlicenseRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest
func GetAlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest() *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest {
	return poolAlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest.Get().(*AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest 将 AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest(v *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest.Put(v)
}
