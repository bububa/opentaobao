package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim 群聊天记录导出接口 API请求
taobao.openim.tribelogs.get

获取openim账号的群聊天记录
*/
type TaobaoOpenimTribelogsGetRequest struct {
    model.Params
    // 群号
    tribeId   string
    // 查询起始时间，UTC秒数。必须在一个月内。
    begin   int64
    // 查询结束时间，UTC秒数。必须大于起始时间并小于当前时间
    end   int64
    // 查询条数
    count   int64
    // 迭代key
    next   string
}

// 初始化TaobaoOpenimTribelogsGetRequest对象
func NewTaobaoOpenimTribelogsGetRequest() *TaobaoOpenimTribelogsGetRequest{
    return &TaobaoOpenimTribelogsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribelogsGetRequest) GetApiMethodName() string {
    return "taobao.openim.tribelogs.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribelogsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TribeId Setter
// 群号
func (r *TaobaoOpenimTribelogsGetRequest) SetTribeId(tribeId string) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribelogsGetRequest) GetTribeId() string {
    return r.tribeId
}
// Begin Setter
// 查询起始时间，UTC秒数。必须在一个月内。
func (r *TaobaoOpenimTribelogsGetRequest) SetBegin(begin int64) error {
    r.begin = begin
    r.Set("begin", begin)
    return nil
}

// Begin Getter
func (r TaobaoOpenimTribelogsGetRequest) GetBegin() int64 {
    return r.begin
}
// End Setter
// 查询结束时间，UTC秒数。必须大于起始时间并小于当前时间
func (r *TaobaoOpenimTribelogsGetRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

// End Getter
func (r TaobaoOpenimTribelogsGetRequest) GetEnd() int64 {
    return r.end
}
// Count Setter
// 查询条数
func (r *TaobaoOpenimTribelogsGetRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

// Count Getter
func (r TaobaoOpenimTribelogsGetRequest) GetCount() int64 {
    return r.count
}
// Next Setter
// 迭代key
func (r *TaobaoOpenimTribelogsGetRequest) SetNext(next string) error {
    r.next = next
    r.Set("next", next)
    return nil
}

// Next Getter
func (r TaobaoOpenimTribelogsGetRequest) GetNext() string {
    return r.next
}
