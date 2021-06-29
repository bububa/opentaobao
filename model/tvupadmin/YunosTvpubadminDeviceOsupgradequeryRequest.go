package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系统升级查询 APIRequest
yunos.tvpubadmin.device.osupgradequery

系统升级查询
*/
type YunosTvpubadminDeviceOsupgradequeryRequest struct {
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

func NewYunosTvpubadminDeviceOsupgradequeryRequest() *YunosTvpubadminDeviceOsupgradequeryRequest{
    return &YunosTvpubadminDeviceOsupgradequeryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.osupgradequery"
}

func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetStatus() string {
    return r.status
}

func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetDayRange(dayRange int64) error {
    r.dayRange = dayRange
    r.Set("day_range", dayRange)
    return nil
}

func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetDayRange() int64 {
    return r.dayRange
}

func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminDeviceOsupgradequeryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminDeviceOsupgradequeryRequest) GetPageSize() int64 {
    return r.pageSize
}

