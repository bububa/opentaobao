package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindevicemodelsAPIRequest 获取品牌下设备列表 API请求
// yunos.tvpubadmin.device.models
//
// 获取品牌下设备列表
type YunostvpubadmindevicemodelsAPIRequest struct {
	model.Params
	// 终端类型
	_terminalType string
	// 牌照方
	_license int64
	// 品牌ID
	_brandId int64
}

// NewYunostvpubadmindevicemodelsRequest 初始化YunostvpubadmindevicemodelsAPIRequest对象
func NewYunostvpubadmindevicemodelsRequest() *YunostvpubadmindevicemodelsAPIRequest {
	return &YunostvpubadmindevicemodelsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindevicemodelsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.models"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindevicemodelsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindevicemodelsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTerminalType is TerminalType Setter
// 终端类型
func (r *YunostvpubadmindevicemodelsAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r YunostvpubadmindevicemodelsAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindevicemodelsAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindevicemodelsAPIRequest) GetLicense() int64 {
	return r._license
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *YunostvpubadmindevicemodelsAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r YunostvpubadmindevicemodelsAPIRequest) GetBrandId() int64 {
	return r._brandId
}
