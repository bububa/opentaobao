package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceQueryAPIRequest 获取设备列表 API请求
// yunos.tvpubadmin.device.query
//
// 获取设备列表
type YunosTvpubadminDeviceQueryAPIRequest struct {
	model.Params
	// 终端类型
	_terminalType string
	// 品牌ID
	_brandId int64
	// 牌照方
	_license int64
	// 页码值
	_pageNo int64
	// 每页条数
	_pageSize int64
}

// NewYunosTvpubadminDeviceQueryRequest 初始化YunosTvpubadminDeviceQueryAPIRequest对象
func NewYunosTvpubadminDeviceQueryRequest() *YunosTvpubadminDeviceQueryAPIRequest {
	return &YunosTvpubadminDeviceQueryAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceQueryAPIRequest) Reset() {
	r._terminalType = ""
	r._brandId = 0
	r._license = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceQueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTerminalType is TerminalType Setter
// 终端类型
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetLicense() int64 {
	return r._license
}

// SetPageNo is PageNo Setter
// 页码值
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolYunosTvpubadminDeviceQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceQueryRequest()
	},
}

// GetYunosTvpubadminDeviceQueryRequest 从 sync.Pool 获取 YunosTvpubadminDeviceQueryAPIRequest
func GetYunosTvpubadminDeviceQueryAPIRequest() *YunosTvpubadminDeviceQueryAPIRequest {
	return poolYunosTvpubadminDeviceQueryAPIRequest.Get().(*YunosTvpubadminDeviceQueryAPIRequest)
}

// ReleaseYunosTvpubadminDeviceQueryAPIRequest 将 YunosTvpubadminDeviceQueryAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceQueryAPIRequest(v *YunosTvpubadminDeviceQueryAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceQueryAPIRequest.Put(v)
}
