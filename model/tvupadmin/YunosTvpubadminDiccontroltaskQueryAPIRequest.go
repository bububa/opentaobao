package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindiccontroltaskqueryAPIRequest 停开服任务列表 API请求
// yunos.tvpubadmin.diccontroltask.query
//
// 牌照方对终端设备的停开服管理
type YunostvpubadmindiccontroltaskqueryAPIRequest struct {
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

// NewYunostvpubadmindiccontroltaskqueryRequest 初始化YunostvpubadmindiccontroltaskqueryAPIRequest对象
func NewYunostvpubadmindiccontroltaskqueryRequest() *YunostvpubadmindiccontroltaskqueryAPIRequest {
	return &YunostvpubadmindiccontroltaskqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindiccontroltaskqueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.diccontroltask.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindiccontroltaskqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindiccontroltaskqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 任务名称
func (r *YunostvpubadmindiccontroltaskqueryAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunostvpubadmindiccontroltaskqueryAPIRequest) GetName() string {
	return r._name
}

// SetStatus is Status Setter
// 任务状态
func (r *YunostvpubadmindiccontroltaskqueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunostvpubadmindiccontroltaskqueryAPIRequest) GetStatus() int64 {
	return r._status
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindiccontroltaskqueryAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindiccontroltaskqueryAPIRequest) GetLicense() int64 {
	return r._license
}

// SetPageNo is PageNo Setter
// 当前页码值
func (r *YunostvpubadmindiccontroltaskqueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunostvpubadmindiccontroltaskqueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *YunostvpubadmindiccontroltaskqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadmindiccontroltaskqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
