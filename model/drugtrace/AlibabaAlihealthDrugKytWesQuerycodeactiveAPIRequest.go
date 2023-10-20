package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest 查询码是否激活 API请求
// alibaba.alihealth.drug.kyt.wes.querycodeactive
//
// 查询码是否激活
type AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 服务时间
	_licenseToken string
	// 码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes string
}

// NewAlibabaAlihealthDrugKytWesQuerycodeactiveRequest 初始化AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest对象
func NewAlibabaAlihealthDrugKytWesQuerycodeactiveRequest() *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest {
	return &AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) Reset() {
	r._refEntId = ""
	r._licenseToken = ""
	r._codes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.querycodeactive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetCodes is Codes Setter
// 码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) SetCodes(_codes string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) GetCodes() string {
	return r._codes
}

var poolAlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytWesQuerycodeactiveRequest()
	},
}

// GetAlibabaAlihealthDrugKytWesQuerycodeactiveRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest
func GetAlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest() *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest {
	return poolAlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest.Get().(*AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest 将 AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest(v *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest.Put(v)
}
