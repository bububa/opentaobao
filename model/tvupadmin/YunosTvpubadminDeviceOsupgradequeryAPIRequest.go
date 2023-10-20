package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceOsupgradequeryAPIRequest 系统升级查询 API请求
// yunos.tvpubadmin.device.osupgradequery
//
// 系统升级查询
type YunosTvpubadminDeviceOsupgradequeryAPIRequest struct {
	model.Params
	// 审核状态
	_status string
	// 牌照方
	_license int64
	// 时间范围
	_dayRange int64
	// 第几页
	_pageNo int64
	// 数据大小
	_pageSize int64
}

// NewYunosTvpubadminDeviceOsupgradequeryRequest 初始化YunosTvpubadminDeviceOsupgradequeryAPIRequest对象
func NewYunosTvpubadminDeviceOsupgradequeryRequest() *YunosTvpubadminDeviceOsupgradequeryAPIRequest {
	return &YunosTvpubadminDeviceOsupgradequeryAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceOsupgradequeryAPIRequest) Reset() {
	r._status = ""
	r._license = 0
	r._dayRange = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceOsupgradequeryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.osupgradequery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceOsupgradequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceOsupgradequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceOsupgradequeryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunosTvpubadminDeviceOsupgradequeryAPIRequest) GetStatus() string {
	return r._status
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceOsupgradequeryAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDeviceOsupgradequeryAPIRequest) GetLicense() int64 {
	return r._license
}

// SetDayRange is DayRange Setter
// 时间范围
func (r *YunosTvpubadminDeviceOsupgradequeryAPIRequest) SetDayRange(_dayRange int64) error {
	r._dayRange = _dayRange
	r.Set("day_range", _dayRange)
	return nil
}

// GetDayRange DayRange Getter
func (r YunosTvpubadminDeviceOsupgradequeryAPIRequest) GetDayRange() int64 {
	return r._dayRange
}

// SetPageNo is PageNo Setter
// 第几页
func (r *YunosTvpubadminDeviceOsupgradequeryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunosTvpubadminDeviceOsupgradequeryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 数据大小
func (r *YunosTvpubadminDeviceOsupgradequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosTvpubadminDeviceOsupgradequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolYunosTvpubadminDeviceOsupgradequeryAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceOsupgradequeryRequest()
	},
}

// GetYunosTvpubadminDeviceOsupgradequeryRequest 从 sync.Pool 获取 YunosTvpubadminDeviceOsupgradequeryAPIRequest
func GetYunosTvpubadminDeviceOsupgradequeryAPIRequest() *YunosTvpubadminDeviceOsupgradequeryAPIRequest {
	return poolYunosTvpubadminDeviceOsupgradequeryAPIRequest.Get().(*YunosTvpubadminDeviceOsupgradequeryAPIRequest)
}

// ReleaseYunosTvpubadminDeviceOsupgradequeryAPIRequest 将 YunosTvpubadminDeviceOsupgradequeryAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceOsupgradequeryAPIRequest(v *YunosTvpubadminDeviceOsupgradequeryAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceOsupgradequeryAPIRequest.Put(v)
}
