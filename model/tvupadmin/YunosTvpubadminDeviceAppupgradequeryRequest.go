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
    license   int64
    // 审核状态
    status   string
    // 时间范围
    dayRange   int64
    // 第几页
    pageNo   int64
    // 数据大小
    pageSize   int64
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
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetLicense() int64 {
    return r.license
}
// Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetStatus() string {
    return r.status
}
// DayRange Setter
// 时间范围
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetDayRange(dayRange int64) error {
    r.dayRange = dayRange
    r.Set("day_range", dayRange)
    return nil
}

// DayRange Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetDayRange() int64 {
    return r.dayRange
}
// PageNo Setter
// 第几页
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 数据大小
func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetPageSize() int64 {
    return r.pageSize
}
