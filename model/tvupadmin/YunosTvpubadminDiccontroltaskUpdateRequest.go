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
type YunosTvpubadminDiccontroltaskUpdateAPIRequest struct {
    model.Params
    // 任务ID
    _id   int64
    // 任务状态
    _status   int64
    // 牌照方
    _license   int64
}

// 初始化YunosTvpubadminDiccontroltaskUpdateAPIRequest对象
func NewYunosTvpubadminDiccontroltaskUpdateRequest() *YunosTvpubadminDiccontroltaskUpdateAPIRequest{
    return &YunosTvpubadminDiccontroltaskUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.diccontroltask.update"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 任务ID
func (r *YunosTvpubadminDiccontroltaskUpdateAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetId() int64 {
    return r._id
}
// Status Setter
// 任务状态
func (r *YunosTvpubadminDiccontroltaskUpdateAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetStatus() int64 {
    return r._status
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDiccontroltaskUpdateAPIRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetLicense() int64 {
    return r._license
}
