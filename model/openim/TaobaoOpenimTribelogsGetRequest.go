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
    _tribeId   string
    // 查询起始时间，UTC秒数。必须在一个月内。
    _begin   int64
    // 查询结束时间，UTC秒数。必须大于起始时间并小于当前时间
    _end   int64
    // 查询条数
    _count   int64
    // 迭代key
    _next   string
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
func (r *TaobaoOpenimTribelogsGetRequest) SetTribeId(_tribeId string) error {
    r._tribeId = _tribeId
    r.Set("tribe_id", _tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribelogsGetRequest) GetTribeId() string {
    return r._tribeId
}
// Begin Setter
// 查询起始时间，UTC秒数。必须在一个月内。
func (r *TaobaoOpenimTribelogsGetRequest) SetBegin(_begin int64) error {
    r._begin = _begin
    r.Set("begin", _begin)
    return nil
}

// Begin Getter
func (r TaobaoOpenimTribelogsGetRequest) GetBegin() int64 {
    return r._begin
}
// End Setter
// 查询结束时间，UTC秒数。必须大于起始时间并小于当前时间
func (r *TaobaoOpenimTribelogsGetRequest) SetEnd(_end int64) error {
    r._end = _end
    r.Set("end", _end)
    return nil
}

// End Getter
func (r TaobaoOpenimTribelogsGetRequest) GetEnd() int64 {
    return r._end
}
// Count Setter
// 查询条数
func (r *TaobaoOpenimTribelogsGetRequest) SetCount(_count int64) error {
    r._count = _count
    r.Set("count", _count)
    return nil
}

// Count Getter
func (r TaobaoOpenimTribelogsGetRequest) GetCount() int64 {
    return r._count
}
// Next Setter
// 迭代key
func (r *TaobaoOpenimTribelogsGetRequest) SetNext(_next string) error {
    r._next = _next
    r.Set("next", _next)
    return nil
}

// Next Getter
func (r TaobaoOpenimTribelogsGetRequest) GetNext() string {
    return r._next
}
