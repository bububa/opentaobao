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
    uid   int64
    // 页码值
    pageNo   int64
    // 每页行数
    pageSize   int64
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
func (r *YunosTvpubadminUserRightsRequest) SetUid(uid int64) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

// Uid Getter
func (r YunosTvpubadminUserRightsRequest) GetUid() int64 {
    return r.uid
}
// PageNo Setter
// 页码值
func (r *YunosTvpubadminUserRightsRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminUserRightsRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页行数
func (r *YunosTvpubadminUserRightsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminUserRightsRequest) GetPageSize() int64 {
    return r.pageSize
}
