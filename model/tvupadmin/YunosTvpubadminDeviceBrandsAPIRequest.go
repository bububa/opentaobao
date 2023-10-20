package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindevicebrandsAPIRequest 获取终端类型下品牌列表 API请求
// yunos.tvpubadmin.device.brands
//
// 获取终端类型下品牌列表
type YunostvpubadmindevicebrandsAPIRequest struct {
	model.Params
	// 终端类型
	_terminalType string
	// 牌照方
	_license int64
}

// NewYunostvpubadmindevicebrandsRequest 初始化YunostvpubadmindevicebrandsAPIRequest对象
func NewYunostvpubadmindevicebrandsRequest() *YunostvpubadmindevicebrandsAPIRequest {
	return &YunostvpubadmindevicebrandsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindevicebrandsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.brands"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindevicebrandsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindevicebrandsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTerminalType is TerminalType Setter
// 终端类型
func (r *YunostvpubadmindevicebrandsAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r YunostvpubadmindevicebrandsAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindevicebrandsAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindevicebrandsAPIRequest) GetLicense() int64 {
	return r._license
}
