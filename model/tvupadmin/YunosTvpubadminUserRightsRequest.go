package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户权益 APIRequest
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

func NewYunosTvpubadminUserRightsRequest() *YunosTvpubadminUserRightsRequest{
    return &YunosTvpubadminUserRightsRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminUserRightsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.user.rights"
}

func (r YunosTvpubadminUserRightsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminUserRightsRequest) SetUid(uid int64) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

func (r YunosTvpubadminUserRightsRequest) GetUid() int64 {
    return r.uid
}

func (r *YunosTvpubadminUserRightsRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminUserRightsRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminUserRightsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminUserRightsRequest) GetPageSize() int64 {
    return r.pageSize
}

