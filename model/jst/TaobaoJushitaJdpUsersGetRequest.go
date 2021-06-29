package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取开通的订单同步服务的用户 API请求
taobao.jushita.jdp.users.get

获取开通的订单同步服务的用户，含有rds的路由关系
*/
type TaobaoJushitaJdpUsersGetRequest struct {
    model.Params
    // 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
    _startModified   string
    // 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
    _endModified   string
    // 每页记录数，默认200
    _pageSize   int64
    // 当前页数
    _pageNo   int64
    // 如果传了user_id表示单条查询
    _userId   int64
}

// 初始化TaobaoJushitaJdpUsersGetRequest对象
func NewTaobaoJushitaJdpUsersGetRequest() *TaobaoJushitaJdpUsersGetRequest{
    return &TaobaoJushitaJdpUsersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJdpUsersGetRequest) GetApiMethodName() string {
    return "taobao.jushita.jdp.users.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJdpUsersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartModified Setter
// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
func (r *TaobaoJushitaJdpUsersGetRequest) SetStartModified(_startModified string) error {
    r._startModified = _startModified
    r.Set("start_modified", _startModified)
    return nil
}

// StartModified Getter
func (r TaobaoJushitaJdpUsersGetRequest) GetStartModified() string {
    return r._startModified
}
// EndModified Setter
// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
func (r *TaobaoJushitaJdpUsersGetRequest) SetEndModified(_endModified string) error {
    r._endModified = _endModified
    r.Set("end_modified", _endModified)
    return nil
}

// EndModified Getter
func (r TaobaoJushitaJdpUsersGetRequest) GetEndModified() string {
    return r._endModified
}
// PageSize Setter
// 每页记录数，默认200
func (r *TaobaoJushitaJdpUsersGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoJushitaJdpUsersGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 当前页数
func (r *TaobaoJushitaJdpUsersGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoJushitaJdpUsersGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// UserId Setter
// 如果传了user_id表示单条查询
func (r *TaobaoJushitaJdpUsersGetRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoJushitaJdpUsersGetRequest) GetUserId() int64 {
    return r._userId
}
