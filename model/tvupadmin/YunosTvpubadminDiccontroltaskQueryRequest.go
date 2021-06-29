package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
停开服任务列表 API请求
yunos.tvpubadmin.diccontroltask.query

牌照方对终端设备的停开服管理
*/
type YunosTvpubadminDiccontroltaskQueryRequest struct {
    model.Params
    // 任务名称
    _name   string
    // 任务状态
    _status   int64
    // 牌照方
    _license   int64
    // 当前页码值
    _pageNo   int64
    // 每页条数
    _pageSize   int64
}

// 初始化YunosTvpubadminDiccontroltaskQueryRequest对象
func NewYunosTvpubadminDiccontroltaskQueryRequest() *YunosTvpubadminDiccontroltaskQueryRequest{
    return &YunosTvpubadminDiccontroltaskQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.diccontroltask.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 任务名称
func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r YunosTvpubadminDiccontroltaskQueryRequest) GetName() string {
    return r._name
}
// Status Setter
// 任务状态
func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r YunosTvpubadminDiccontroltaskQueryRequest) GetStatus() int64 {
    return r._status
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDiccontroltaskQueryRequest) GetLicense() int64 {
    return r._license
}
// PageNo Setter
// 当前页码值
func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminDiccontroltaskQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页条数
func (r *YunosTvpubadminDiccontroltaskQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDiccontroltaskQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
