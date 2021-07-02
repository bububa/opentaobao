package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceModelsAPIRequest 获取品牌下设备列表 API请求
// yunos.tvpubadmin.device.models
//
// 获取品牌下设备列表
type YunosTvpubadminDeviceModelsAPIRequest struct {
	model.Params
	// 终端类型
	_terminalType string
	// 品牌ID
	_brandId int64
	// 牌照方
	_license int64
}

// NewYunosTvpubadminDeviceModelsRequest 初始化YunosTvpubadminDeviceModelsAPIRequest对象
func NewYunosTvpubadminDeviceModelsRequest() *YunosTvpubadminDeviceModelsAPIRequest {
	return &YunosTvpubadminDeviceModelsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceModelsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.models"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceModelsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTerminalType is TerminalType Setter
// 终端类型
func (r *YunosTvpubadminDeviceModelsAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r YunosTvpubadminDeviceModelsAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *YunosTvpubadminDeviceModelsAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r YunosTvpubadminDeviceModelsAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceModelsAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDeviceModelsAPIRequest) GetLicense() int64 {
	return r._license
}
