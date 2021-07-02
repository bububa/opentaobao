package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDiccontroltaskQueryAPIRequest 停开服任务列表 API请求
// yunos.tvpubadmin.diccontroltask.query
//
// 牌照方对终端设备的停开服管理
type YunosTvpubadminDiccontroltaskQueryAPIRequest struct {
	model.Params
	// 任务名称
	_name string
	// 任务状态
	_status int64
	// 牌照方
	_license int64
	// 当前页码值
	_pageNo int64
	// 每页条数
	_pageSize int64
}

// NewYunosTvpubadminDiccontroltaskQueryRequest 初始化YunosTvpubadminDiccontroltaskQueryAPIRequest对象
func NewYunosTvpubadminDiccontroltaskQueryRequest() *YunosTvpubadminDiccontroltaskQueryAPIRequest {
	return &YunosTvpubadminDiccontroltaskQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskQueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.diccontroltask.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetName is Name Setter
// 任务名称
func (r *YunosTvpubadminDiccontroltaskQueryAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunosTvpubadminDiccontroltaskQueryAPIRequest) GetName() string {
	return r._name
}

// SetStatus is Status Setter
// 任务状态
func (r *YunosTvpubadminDiccontroltaskQueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunosTvpubadminDiccontroltaskQueryAPIRequest) GetStatus() int64 {
	return r._status
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDiccontroltaskQueryAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDiccontroltaskQueryAPIRequest) GetLicense() int64 {
	return r._license
}

// SetPageNo is PageNo Setter
// 当前页码值
func (r *YunosTvpubadminDiccontroltaskQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunosTvpubadminDiccontroltaskQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *YunosTvpubadminDiccontroltaskQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosTvpubadminDiccontroltaskQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
