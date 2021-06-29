package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
停开服任务状态变更 APIRequest
yunos.tvpubadmin.diccontroltask.update

停开服任务状态变更
*/
type YunosTvpubadminDiccontroltaskUpdateRequest struct {
    model.Params

    // 任务ID
    id   int64 

    // 任务状态
    status   int64 

    // 牌照方
    license   int64 

}

func NewYunosTvpubadminDiccontroltaskUpdateRequest() *YunosTvpubadminDiccontroltaskUpdateRequest{
    return &YunosTvpubadminDiccontroltaskUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.diccontroltask.update"
}

func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDiccontroltaskUpdateRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetId() int64 {
    return r.id
}

func (r *YunosTvpubadminDiccontroltaskUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetStatus() int64 {
    return r.status
}

func (r *YunosTvpubadminDiccontroltaskUpdateRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetLicense() int64 {
    return r.license
}

