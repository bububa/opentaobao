package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceBrandsAPIRequest 获取终端类型下品牌列表 API请求
// yunos.tvpubadmin.device.brands
//
// 获取终端类型下品牌列表
type YunosTvpubadminDeviceBrandsAPIRequest struct {
	model.Params
	// 终端类型
	_terminalType string
	// 牌照方
	_license int64
}

// NewYunosTvpubadminDeviceBrandsRequest 初始化YunosTvpubadminDeviceBrandsAPIRequest对象
func NewYunosTvpubadminDeviceBrandsRequest() *YunosTvpubadminDeviceBrandsAPIRequest {
	return &YunosTvpubadminDeviceBrandsAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceBrandsAPIRequest) Reset() {
	r._terminalType = ""
	r._license = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceBrandsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.brands"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceBrandsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceBrandsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTerminalType is TerminalType Setter
// 终端类型
func (r *YunosTvpubadminDeviceBrandsAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r YunosTvpubadminDeviceBrandsAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceBrandsAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDeviceBrandsAPIRequest) GetLicense() int64 {
	return r._license
}

var poolYunosTvpubadminDeviceBrandsAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceBrandsRequest()
	},
}

// GetYunosTvpubadminDeviceBrandsRequest 从 sync.Pool 获取 YunosTvpubadminDeviceBrandsAPIRequest
func GetYunosTvpubadminDeviceBrandsAPIRequest() *YunosTvpubadminDeviceBrandsAPIRequest {
	return poolYunosTvpubadminDeviceBrandsAPIRequest.Get().(*YunosTvpubadminDeviceBrandsAPIRequest)
}

// ReleaseYunosTvpubadminDeviceBrandsAPIRequest 将 YunosTvpubadminDeviceBrandsAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceBrandsAPIRequest(v *YunosTvpubadminDeviceBrandsAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceBrandsAPIRequest.Put(v)
}
