package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentdevicegetvendorAPIRequest 查询设备Vendor信息 API请求
// yunos.tvpubadmin.content.device.getvendor
//
// 查询设备Vendor信息
type YunostvpubadmincontentdevicegetvendorAPIRequest struct {
	model.Params
	// license
	_license int64
	// brand_id
	_brandId int64
}

// NewYunostvpubadmincontentdevicegetvendorRequest 初始化YunostvpubadmincontentdevicegetvendorAPIRequest对象
func NewYunostvpubadmincontentdevicegetvendorRequest() *YunostvpubadmincontentdevicegetvendorAPIRequest {
	return &YunostvpubadmincontentdevicegetvendorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentdevicegetvendorAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.device.getvendor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentdevicegetvendorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentdevicegetvendorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLicense is License Setter
// license
func (r *YunostvpubadmincontentdevicegetvendorAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmincontentdevicegetvendorAPIRequest) GetLicense() int64 {
	return r._license
}

// SetBrandId is BrandId Setter
// brand_id
func (r *YunostvpubadmincontentdevicegetvendorAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r YunostvpubadmincontentdevicegetvendorAPIRequest) GetBrandId() int64 {
	return r._brandId
}
