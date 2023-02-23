package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentDeviceGetvendorAPIRequest 查询设备Vendor信息 API请求
// yunos.tvpubadmin.content.device.getvendor
//
// 查询设备Vendor信息
type YunosTvpubadminContentDeviceGetvendorAPIRequest struct {
	model.Params
	// license
	_license int64
	// brand_id
	_brandId int64
}

// NewYunosTvpubadminContentDeviceGetvendorRequest 初始化YunosTvpubadminContentDeviceGetvendorAPIRequest对象
func NewYunosTvpubadminContentDeviceGetvendorRequest() *YunosTvpubadminContentDeviceGetvendorAPIRequest {
	return &YunosTvpubadminContentDeviceGetvendorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentDeviceGetvendorAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.device.getvendor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentDeviceGetvendorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentDeviceGetvendorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLicense is License Setter
// license
func (r *YunosTvpubadminContentDeviceGetvendorAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminContentDeviceGetvendorAPIRequest) GetLicense() int64 {
	return r._license
}

// SetBrandId is BrandId Setter
// brand_id
func (r *YunosTvpubadminContentDeviceGetvendorAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r YunosTvpubadminContentDeviceGetvendorAPIRequest) GetBrandId() int64 {
	return r._brandId
}
