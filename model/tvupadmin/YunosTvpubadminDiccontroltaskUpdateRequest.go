package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
停开服任务状态变更 API请求
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

// 初始化YunosTvpubadminDiccontroltaskUpdateRequest对象
func NewYunosTvpubadminDiccontroltaskUpdateRequest() *YunosTvpubadminDiccontroltaskUpdateRequest{
    return &YunosTvpubadminDiccontroltaskUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.diccontroltask.update"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 任务ID
func (r *YunosTvpubadminDiccontroltaskUpdateRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetId() int64 {
    return r.id
}
// Status Setter
// 任务状态
func (r *YunosTvpubadminDiccontroltaskUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetStatus() int64 {
    return r.status
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDiccontroltaskUpdateRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminDiccontroltaskUpdateRequest) GetLicense() int64 {
    return r.license
}
