package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesGetbyentidAPIRequest 根据企业主键查看企业详细信息 API请求
// alibaba.alihealth.drug.kyt.wes.getbyentid
//
// 根据企业主键查看企业详细信息
type AlibabaAlihealthDrugKytWesGetbyentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 服务时间
	_licenseToken string
	// 准备要查询的企业ID（返回该企业ID的详细信息）
	_entId string
}

// NewAlibabaAlihealthDrugKytWesGetbyentidRequest 初始化AlibabaAlihealthDrugKytWesGetbyentidAPIRequest对象
func NewAlibabaAlihealthDrugKytWesGetbyentidRequest() *AlibabaAlihealthDrugKytWesGetbyentidAPIRequest {
	return &AlibabaAlihealthDrugKytWesGetbyentidAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) Reset() {
	r._refEntId = ""
	r._licenseToken = ""
	r._entId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.getbyentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetEntId is EntId Setter
// 准备要查询的企业ID（返回该企业ID的详细信息）
func (r *AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) SetEntId(_entId string) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// GetEntId EntId Getter
func (r AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) GetEntId() string {
	return r._entId
}

var poolAlibabaAlihealthDrugKytWesGetbyentidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytWesGetbyentidRequest()
	},
}

// GetAlibabaAlihealthDrugKytWesGetbyentidRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesGetbyentidAPIRequest
func GetAlibabaAlihealthDrugKytWesGetbyentidAPIRequest() *AlibabaAlihealthDrugKytWesGetbyentidAPIRequest {
	return poolAlibabaAlihealthDrugKytWesGetbyentidAPIRequest.Get().(*AlibabaAlihealthDrugKytWesGetbyentidAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytWesGetbyentidAPIRequest 将 AlibabaAlihealthDrugKytWesGetbyentidAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesGetbyentidAPIRequest(v *AlibabaAlihealthDrugKytWesGetbyentidAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesGetbyentidAPIRequest.Put(v)
}
