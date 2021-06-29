package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取openim账号的聊天关系 API请求
taobao.openim.relations.get

获取openim账号的聊天关系
*/
type TaobaoOpenimRelationsGetRequest struct {
    model.Params
    // 查询起始日期。格式yyyyMMdd。不得早于一个月
    begDate   string
    // 查询结束日期。格式yyyyMMdd。不得早于一个月
    endDate   string
    // 用户信息
    user   *OpenImUser
}

// 初始化TaobaoOpenimRelationsGetRequest对象
func NewTaobaoOpenimRelationsGetRequest() *TaobaoOpenimRelationsGetRequest{
    return &TaobaoOpenimRelationsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimRelationsGetRequest) GetApiMethodName() string {
    return "taobao.openim.relations.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimRelationsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BegDate Setter
// 查询起始日期。格式yyyyMMdd。不得早于一个月
func (r *TaobaoOpenimRelationsGetRequest) SetBegDate(begDate string) error {
    r.begDate = begDate
    r.Set("beg_date", begDate)
    return nil
}

// BegDate Getter
func (r TaobaoOpenimRelationsGetRequest) GetBegDate() string {
    return r.begDate
}
// EndDate Setter
// 查询结束日期。格式yyyyMMdd。不得早于一个月
func (r *TaobaoOpenimRelationsGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoOpenimRelationsGetRequest) GetEndDate() string {
    return r.endDate
}
// User Setter
// 用户信息
func (r *TaobaoOpenimRelationsGetRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimRelationsGetRequest) GetUser() *OpenImUser {
    return r.user
}
