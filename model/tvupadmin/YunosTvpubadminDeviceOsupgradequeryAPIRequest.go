package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceosupgradequeryAPIRequest 系统升级查询 API请求
// yunos.tvpubadmin.device.osupgradequery
//
// 系统升级查询
type YunostvpubadmindeviceosupgradequeryAPIRequest struct {
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

// NewYunostvpubadmindeviceosupgradequeryRequest 初始化YunostvpubadmindeviceosupgradequeryAPIRequest对象
func NewYunostvpubadmindeviceosupgradequeryRequest() *YunostvpubadmindeviceosupgradequeryAPIRequest {
	return &YunostvpubadmindeviceosupgradequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceosupgradequeryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.osupgradequery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceosupgradequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceosupgradequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 审核状态
func (r *YunostvpubadmindeviceosupgradequeryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunostvpubadmindeviceosupgradequeryAPIRequest) GetStatus() string {
	return r._status
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindeviceosupgradequeryAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindeviceosupgradequeryAPIRequest) GetLicense() int64 {
	return r._license
}

// SetDayRange is DayRange Setter
// 时间范围
func (r *YunostvpubadmindeviceosupgradequeryAPIRequest) SetDayRange(_dayRange int64) error {
	r._dayRange = _dayRange
	r.Set("day_range", _dayRange)
	return nil
}

// GetDayRange DayRange Getter
func (r YunostvpubadmindeviceosupgradequeryAPIRequest) GetDayRange() int64 {
	return r._dayRange
}

// SetPageNo is PageNo Setter
// 第几页
func (r *YunostvpubadmindeviceosupgradequeryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunostvpubadmindeviceosupgradequeryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 数据大小
func (r *YunostvpubadmindeviceosupgradequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadmindeviceosupgradequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
