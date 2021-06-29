package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim聊天记录查询接口 API请求
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

// 初始化TaobaoOpenimChatlogsGetRequest对象
func NewTaobaoOpenimChatlogsGetRequest() *TaobaoOpenimChatlogsGetRequest{
    return &TaobaoOpenimChatlogsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimChatlogsGetRequest) GetApiMethodName() string {
    return "taobao.openim.chatlogs.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimChatlogsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User1 Setter
// 用户1信息
func (r *TaobaoOpenimChatlogsGetRequest) SetUser1(user1 *OpenImUser) error {
    r.user1 = user1
    r.Set("user1", user1)
    return nil
}

// User1 Getter
func (r TaobaoOpenimChatlogsGetRequest) GetUser1() *OpenImUser {
    return r.user1
}
// User2 Setter
// 用户2信息
func (r *TaobaoOpenimChatlogsGetRequest) SetUser2(user2 *OpenImUser) error {
    r.user2 = user2
    r.Set("user2", user2)
    return nil
}

// User2 Getter
func (r TaobaoOpenimChatlogsGetRequest) GetUser2() *OpenImUser {
    return r.user2
}
// Begin Setter
// 查询开始时间（UTC时间）
func (r *TaobaoOpenimChatlogsGetRequest) SetBegin(begin int64) error {
    r.begin = begin
    r.Set("begin", begin)
    return nil
}

// Begin Getter
func (r TaobaoOpenimChatlogsGetRequest) GetBegin() int64 {
    return r.begin
}
// End Setter
// 查询结束时间（UTC时间）
func (r *TaobaoOpenimChatlogsGetRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

// End Getter
func (r TaobaoOpenimChatlogsGetRequest) GetEnd() int64 {
    return r.end
}
// Count Setter
// 查询条数
func (r *TaobaoOpenimChatlogsGetRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

// Count Getter
func (r TaobaoOpenimChatlogsGetRequest) GetCount() int64 {
    return r.count
}
// NextKey Setter
// 迭代key
func (r *TaobaoOpenimChatlogsGetRequest) SetNextKey(nextKey string) error {
    r.nextKey = nextKey
    r.Set("next_key", nextKey)
    return nil
}

// NextKey Getter
func (r TaobaoOpenimChatlogsGetRequest) GetNextKey() string {
    return r.nextKey
}
