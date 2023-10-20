package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwesgetbyentidAPIRequest 根据企业主键查看企业详细信息 API请求
// alibaba.alihealth.drug.kyt.wes.getbyentid
//
// 根据企业主键查看企业详细信息
type AlibabaalihealthdrugkytwesgetbyentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 服务时间
	_licenseToken string
	// 准备要查询的企业ID（返回该企业ID的详细信息）
	_entId string
}

// NewAlibabaalihealthdrugkytwesgetbyentidRequest 初始化AlibabaalihealthdrugkytwesgetbyentidAPIRequest对象
func NewAlibabaalihealthdrugkytwesgetbyentidRequest() *AlibabaalihealthdrugkytwesgetbyentidAPIRequest {
	return &AlibabaalihealthdrugkytwesgetbyentidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytwesgetbyentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.getbyentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytwesgetbyentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytwesgetbyentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaalihealthdrugkytwesgetbyentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytwesgetbyentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaalihealthdrugkytwesgetbyentidAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaalihealthdrugkytwesgetbyentidAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetEntId is EntId Setter
// 准备要查询的企业ID（返回该企业ID的详细信息）
func (r *AlibabaalihealthdrugkytwesgetbyentidAPIRequest) SetEntId(_entId string) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// GetEntId EntId Getter
func (r AlibabaalihealthdrugkytwesgetbyentidAPIRequest) GetEntId() string {
	return r._entId
}
