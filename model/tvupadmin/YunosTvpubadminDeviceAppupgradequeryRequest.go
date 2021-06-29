package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用升级查询 API请求
yunos.tvpubadmin.device.appupgradequery

应用升级查询
*/
type YunosTvpubadminDeviceAppupgradequeryRequest struct {
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

// 初始化YunosTvpubadminDeviceAppupgradequeryRequest对象
func NewYunosTvpubadminDeviceAppupgradequeryRequest() *YunosTvpubadminDeviceAppupgradequeryRequest{
    return &YunosTvpubadminDeviceAppupgradequeryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.appupgradequery"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetLicense() int64 {
    return r._license
}
// Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetStatus() string {
    return r._status
}
// DayRange Setter
// 时间范围
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetDayRange(_dayRange int64) error {
    r._dayRange = _dayRange
    r.Set("day_range", _dayRange)
    return nil
}

// DayRange Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetDayRange() int64 {
    return r._dayRange
}
// PageNo Setter
// 第几页
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 数据大小
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetPageSize() int64 {
    return r._pageSize
}
