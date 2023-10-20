package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest wes查询追溯码信息 API请求
// alibaba.alihealth.drug.code.kyt.wes.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest struct {
	model.Params
	// 服务时间
	_licenseToken []string
	// 企业唯一标识
	_refEntId string
	// 码列表【多个码用逗号分隔的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes string
}

// NewAlibabaAlihealthDrugCodeKytWesQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytWesQuerycodeRequest() *AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) Reset() {
	r._licenseToken = r._licenseToken[:0]
	r._refEntId = ""
	r._codes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.wes.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) SetLicenseToken(_licenseToken []string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) GetLicenseToken() []string {
	return r._licenseToken
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCodes is Codes Setter
// 码列表【多个码用逗号分隔的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) SetCodes(_codes string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) GetCodes() string {
	return r._codes
}

var poolAlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeKytWesQuerycodeRequest()
	},
}

// GetAlibabaAlihealthDrugCodeKytWesQuerycodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest
func GetAlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest() *AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest {
	return poolAlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest.Get().(*AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest 将 AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest(v *AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest.Put(v)
}
