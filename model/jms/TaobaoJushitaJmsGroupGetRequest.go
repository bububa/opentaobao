package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询ONS分组 API请求
taobao.jushita.jms.group.get

查询当前appkey在ONS中已有的分组
*/
type TaobaoJushitaJmsGroupGetAPIRequest struct {
    model.Params
    // 页码
    _pageNo   int64
    // 每页返回多少个分组
    _pageSize   int64
    // 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
    _groupNames   []string
}

// 初始化TaobaoJushitaJmsGroupGetAPIRequest对象
func NewTaobaoJushitaJmsGroupGetRequest() *TaobaoJushitaJmsGroupGetAPIRequest{
    return &TaobaoJushitaJmsGroupGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.group.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNo Setter
// 页码
func (r *TaobaoJushitaJmsGroupGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页返回多少个分组
func (r *TaobaoJushitaJmsGroupGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// GroupNames Setter
// 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
func (r *TaobaoJushitaJmsGroupGetAPIRequest) SetGroupNames(_groupNames []string) error {
    r._groupNames = _groupNames
    r.Set("group_names", _groupNames)
    return nil
}

// GroupNames Getter
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetGroupNames() []string {
    return r._groupNames
}
