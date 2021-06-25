package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取开通的订单同步服务的用户 APIRequest
taobao.jushita.jdp.users.get

获取开通的订单同步服务的用户，含有rds的路由关系
*/
type TaobaoJushitaJdpUsersGetRequest struct {
    model.Params

    // 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
    startModified   string 

    // 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
    endModified   string 

    // 每页记录数，默认200
    pageSize   int64 

    // 当前页数
    pageNo   int64 

    // 如果传了user_id表示单条查询
    userId   int64 

}

func NewTaobaoJushitaJdpUsersGetRequest() *TaobaoJushitaJdpUsersGetRequest{
    return &TaobaoJushitaJdpUsersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJushitaJdpUsersGetRequest) GetApiMethodName() string {
    return "taobao.jushita.jdp.users.get"
}

func (r TaobaoJushitaJdpUsersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJushitaJdpUsersGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

func (r TaobaoJushitaJdpUsersGetRequest) GetStartModified() string {
    return r.startModified
}

func (r *TaobaoJushitaJdpUsersGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

func (r TaobaoJushitaJdpUsersGetRequest) GetEndModified() string {
    return r.endModified
}

func (r *TaobaoJushitaJdpUsersGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoJushitaJdpUsersGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoJushitaJdpUsersGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoJushitaJdpUsersGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoJushitaJdpUsersGetRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoJushitaJdpUsersGetRequest) GetUserId() int64 {
    return r.userId
}

