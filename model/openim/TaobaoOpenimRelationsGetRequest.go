package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取openim账号的聊天关系 APIRequest
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

func NewTaobaoOpenimRelationsGetRequest() *TaobaoOpenimRelationsGetRequest{
    return &TaobaoOpenimRelationsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimRelationsGetRequest) GetApiMethodName() string {
    return "taobao.openim.relations.get"
}

func (r TaobaoOpenimRelationsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimRelationsGetRequest) SetBegDate(begDate string) error {
    r.begDate = begDate
    r.Set("beg_date", begDate)
    return nil
}

func (r TaobaoOpenimRelationsGetRequest) GetBegDate() string {
    return r.begDate
}

func (r *TaobaoOpenimRelationsGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoOpenimRelationsGetRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoOpenimRelationsGetRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimRelationsGetRequest) GetUser() *OpenImUser {
    return r.user
}

