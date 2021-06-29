package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户权益 API请求
yunos.tvpubadmin.user.rights

获取用户权益
*/
type YunosTvpubadminUserRightsRequest struct {
    model.Params
    // 用户ID
    _uid   int64
    // 页码值
    _pageNo   int64
    // 每页行数
    _pageSize   int64
}

// 初始化YunosTvpubadminUserRightsRequest对象
func NewYunosTvpubadminUserRightsRequest() *YunosTvpubadminUserRightsRequest{
    return &YunosTvpubadminUserRightsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminUserRightsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.user.rights"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminUserRightsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uid Setter
// 用户ID
func (r *YunosTvpubadminUserRightsRequest) SetUid(_uid int64) error {
    r._uid = _uid
    r.Set("uid", _uid)
    return nil
}

// Uid Getter
func (r YunosTvpubadminUserRightsRequest) GetUid() int64 {
    return r._uid
}
// PageNo Setter
// 页码值
func (r *YunosTvpubadminUserRightsRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminUserRightsRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页行数
func (r *YunosTvpubadminUserRightsRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminUserRightsRequest) GetPageSize() int64 {
    return r._pageSize
}
