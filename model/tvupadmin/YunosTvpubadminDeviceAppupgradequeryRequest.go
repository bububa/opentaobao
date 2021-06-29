package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用升级查询 APIRequest
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

func NewYunosTvpubadminDeviceAppupgradequeryRequest() *YunosTvpubadminDeviceAppupgradequeryRequest{
    return &YunosTvpubadminDeviceAppupgradequeryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.appupgradequery"
}

func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetStatus() string {
    return r.status
}

func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetDayRange(dayRange int64) error {
    r.dayRange = dayRange
    r.Set("day_range", dayRange)
    return nil
}

func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetDayRange() int64 {
    return r.dayRange
}

func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminDeviceAppupgradequeryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminDeviceAppupgradequeryRequest) GetPageSize() int64 {
    return r.pageSize
}

