package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系统升级查询 API请求
yunos.tvpubadmin.device.osupgradequery

系统升级查询
*/
type YunosTvpubadminDeviceOsupgradequeryRequest struct {
    model.Params
    // 牌照方
    _license   int64
    // 审核状态
    _status   string
    // 时间范围
    _dayRange   int64
    // 第几页
    _pageNo   int64
    // 数据大小
    _pageSize   int64
}

// 初始化YunosTvpubadminDeviceOsupgradequeryRequest对象
func NewYunosTvpubadminDeviceOsupgradequeryRequest() *YunosTvpubadminDeviceOsupgradequeryRequest{
    return &YunosTvpubadminDeviceOsupgradequeryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.osupgradequery"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetLicense() int64 {
    return r._license
}
// Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetStatus() string {
    return r._status
}
// DayRange Setter
// 时间范围
func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetDayRange(_dayRange int64) error {
    r._dayRange = _dayRange
    r.Set("day_range", _dayRange)
    return nil
}

// DayRange Getter
func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetDayRange() int64 {
    return r._dayRange
}
// PageNo Setter
// 第几页
func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 数据大小
func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetPageSize() int64 {
    return r._pageSize
}
