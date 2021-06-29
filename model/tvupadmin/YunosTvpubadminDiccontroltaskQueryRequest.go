package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
停开服任务列表 APIRequest
yunos.tvpubadmin.diccontroltask.query

牌照方对终端设备的停开服管理
*/
type YunosTvpubadminDiccontroltaskQueryRequest struct {
    model.Params

    // 任务名称
    name   string 

    // 任务状态
    status   int64 

    // 牌照方
    license   int64 

    // 当前页码值
    pageNo   int64 

    // 每页条数
    pageSize   int64 

}

func NewYunosTvpubadminDiccontroltaskQueryRequest() *YunosTvpubadminDiccontroltaskQueryRequest{
    return &YunosTvpubadminDiccontroltaskQueryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDiccontroltaskQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.diccontroltask.query"
}

func (r YunosTvpubadminDiccontroltaskQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r YunosTvpubadminDiccontroltaskQueryRequest) GetName() string {
    return r.name
}

func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosTvpubadminDiccontroltaskQueryRequest) GetStatus() int64 {
    return r.status
}

func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDiccontroltaskQueryRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminDiccontroltaskQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminDiccontroltaskQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

