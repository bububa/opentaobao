package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceappupgradequeryAPIRequest 应用升级查询 API请求
// yunos.tvpubadmin.device.appupgradequery
//
// 应用升级查询
type YunostvpubadmindeviceappupgradequeryAPIRequest struct {
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

// NewYunostvpubadmindeviceappupgradequeryRequest 初始化YunostvpubadmindeviceappupgradequeryAPIRequest对象
func NewYunostvpubadmindeviceappupgradequeryRequest() *YunostvpubadmindeviceappupgradequeryAPIRequest {
	return &YunostvpubadmindeviceappupgradequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceappupgradequeryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.appupgradequery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceappupgradequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceappupgradequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 审核状态
func (r *YunostvpubadmindeviceappupgradequeryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunostvpubadmindeviceappupgradequeryAPIRequest) GetStatus() string {
	return r._status
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindeviceappupgradequeryAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindeviceappupgradequeryAPIRequest) GetLicense() int64 {
	return r._license
}

// SetDayRange is DayRange Setter
// 时间范围
func (r *YunostvpubadmindeviceappupgradequeryAPIRequest) SetDayRange(_dayRange int64) error {
	r._dayRange = _dayRange
	r.Set("day_range", _dayRange)
	return nil
}

// GetDayRange DayRange Getter
func (r YunostvpubadmindeviceappupgradequeryAPIRequest) GetDayRange() int64 {
	return r._dayRange
}

// SetPageNo is PageNo Setter
// 第几页
func (r *YunostvpubadmindeviceappupgradequeryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunostvpubadmindeviceappupgradequeryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 数据大小
func (r *YunostvpubadmindeviceappupgradequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadmindeviceappupgradequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
