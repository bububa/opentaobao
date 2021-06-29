package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取停开服任务详情 API请求
yunos.tvpubadmin.diccontroltask.getinfo

获取停开服任务详情
*/
type YunosTvpubadminDiccontroltaskGetinfoRequest struct {
    model.Params
    // 任务ID
    _id   int64
    // 牌照方
    _license   int64
}

// 初始化YunosTvpubadminDiccontroltaskGetinfoRequest对象
func NewYunosTvpubadminDiccontroltaskGetinfoRequest() *YunosTvpubadminDiccontroltaskGetinfoRequest{
    return &YunosTvpubadminDiccontroltaskGetinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskGetinfoRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.diccontroltask.getinfo"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskGetinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 任务ID
func (r *YunosTvpubadminDiccontroltaskGetinfoRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminDiccontroltaskGetinfoRequest) GetId() int64 {
    return r._id
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDiccontroltaskGetinfoRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDiccontroltaskGetinfoRequest) GetLicense() int64 {
    return r._license
}
