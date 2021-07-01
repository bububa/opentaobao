package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceAppupgradequeryAPIRequest
应用升级查询 API请求
yunos.tvpubadmin.device.appupgradequery

应用升级查询 */
type YunosTvpubadminDeviceAppupgradequeryAPIRequest struct {
	model.Params
	// 牌照方
	_license int64
	// 审核状态
	_status string
	// 时间范围
	_dayRange int64
	// 第几页
	_pageNo int64
	// 数据大小
	_pageSize int64
}

// NewYunosTvpubadminDeviceAppupgradequeryRequest 初始化YunosTvpubadminDeviceAppupgradequeryAPIRequest对象
func NewYunosTvpubadminDeviceAppupgradequeryRequest() *YunosTvpubadminDeviceAppupgradequeryAPIRequest {
	return &YunosTvpubadminDeviceAppupgradequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceAppupgradequeryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.appupgradequery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceAppupgradequeryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceAppupgradequeryAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// Get License Getter
func (r YunosTvpubadminDeviceAppupgradequeryAPIRequest) GetLicense() int64 {
	return r._license
}

// Set is Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceAppupgradequeryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r YunosTvpubadminDeviceAppupgradequeryAPIRequest) GetStatus() string {
	return r._status
}

// Set is DayRange Setter
// 时间范围
func (r *YunosTvpubadminDeviceAppupgradequeryAPIRequest) SetDayRange(_dayRange int64) error {
	r._dayRange = _dayRange
	r.Set("day_range", _dayRange)
	return nil
}

// Get DayRange Getter
func (r YunosTvpubadminDeviceAppupgradequeryAPIRequest) GetDayRange() int64 {
	return r._dayRange
}

// Set is PageNo Setter
// 第几页
func (r *YunosTvpubadminDeviceAppupgradequeryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r YunosTvpubadminDeviceAppupgradequeryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 数据大小
func (r *YunosTvpubadminDeviceAppupgradequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r YunosTvpubadminDeviceAppupgradequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
