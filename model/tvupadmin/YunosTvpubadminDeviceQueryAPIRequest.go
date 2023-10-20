package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindevicequeryAPIRequest 获取设备列表 API请求
// yunos.tvpubadmin.device.query
//
// 获取设备列表
type YunostvpubadmindevicequeryAPIRequest struct {
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

// NewYunostvpubadmindevicequeryRequest 初始化YunostvpubadmindevicequeryAPIRequest对象
func NewYunostvpubadmindevicequeryRequest() *YunostvpubadmindevicequeryAPIRequest {
	return &YunostvpubadmindevicequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindevicequeryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindevicequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindevicequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTerminalType is TerminalType Setter
// 终端类型
func (r *YunostvpubadmindevicequeryAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r YunostvpubadmindevicequeryAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *YunostvpubadmindevicequeryAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r YunostvpubadmindevicequeryAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindevicequeryAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindevicequeryAPIRequest) GetLicense() int64 {
	return r._license
}

// SetPageNo is PageNo Setter
// 页码值
func (r *YunostvpubadmindevicequeryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunostvpubadmindevicequeryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *YunostvpubadmindevicequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadmindevicequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
