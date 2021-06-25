package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim聊天记录查询接口 APIRequest
taobao.openim.chatlogs.get

查询openim账号聊天记录
*/
type TaobaoOpenimChatlogsGetRequest struct {
    model.Params

    // 用户1信息
    user1   *OpenImUser 

    // 用户2信息
    user2   *OpenImUser 

    // 查询开始时间（UTC时间）
    begin   int64 

    // 查询结束时间（UTC时间）
    end   int64 

    // 查询条数
    count   int64 

    // 迭代key
    nextKey   string 

}

func NewTaobaoOpenimChatlogsGetRequest() *TaobaoOpenimChatlogsGetRequest{
    return &TaobaoOpenimChatlogsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimChatlogsGetRequest) GetApiMethodName() string {
    return "taobao.openim.chatlogs.get"
}

func (r TaobaoOpenimChatlogsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimChatlogsGetRequest) SetUser1(user1 *OpenImUser) error {
    r.user1 = user1
    r.Set("user1", user1)
    return nil
}

func (r TaobaoOpenimChatlogsGetRequest) GetUser1() *OpenImUser {
    return r.user1
}

func (r *TaobaoOpenimChatlogsGetRequest) SetUser2(user2 *OpenImUser) error {
    r.user2 = user2
    r.Set("user2", user2)
    return nil
}

func (r TaobaoOpenimChatlogsGetRequest) GetUser2() *OpenImUser {
    return r.user2
}

func (r *TaobaoOpenimChatlogsGetRequest) SetBegin(begin int64) error {
    r.begin = begin
    r.Set("begin", begin)
    return nil
}

func (r TaobaoOpenimChatlogsGetRequest) GetBegin() int64 {
    return r.begin
}

func (r *TaobaoOpenimChatlogsGetRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

func (r TaobaoOpenimChatlogsGetRequest) GetEnd() int64 {
    return r.end
}

func (r *TaobaoOpenimChatlogsGetRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

func (r TaobaoOpenimChatlogsGetRequest) GetCount() int64 {
    return r.count
}

func (r *TaobaoOpenimChatlogsGetRequest) SetNextKey(nextKey string) error {
    r.nextKey = nextKey
    r.Set("next_key", nextKey)
    return nil
}

func (r TaobaoOpenimChatlogsGetRequest) GetNextKey() string {
    return r.nextKey
}

