package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim应用聊天记录查询 API请求
taobao.openim.app.chatlogs.get

查询openim应用的聊天记录
*/
type TaobaoOpenimAppChatlogsGetRequest struct {
    model.Params
    // 查询结束时间。UTC时间。精度到秒
    _beg   int64
    // 查询结束时间。UTC时间。精度到秒
    _end   int64
    // 查询最大条数
    _count   int64
    // 迭代key
    _next   string
}

// 初始化TaobaoOpenimAppChatlogsGetRequest对象
func NewTaobaoOpenimAppChatlogsGetRequest() *TaobaoOpenimAppChatlogsGetRequest{
    return &TaobaoOpenimAppChatlogsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimAppChatlogsGetRequest) GetApiMethodName() string {
    return "taobao.openim.app.chatlogs.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimAppChatlogsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Beg Setter
// 查询结束时间。UTC时间。精度到秒
func (r *TaobaoOpenimAppChatlogsGetRequest) SetBeg(_beg int64) error {
    r._beg = _beg
    r.Set("beg", _beg)
    return nil
}

// Beg Getter
func (r TaobaoOpenimAppChatlogsGetRequest) GetBeg() int64 {
    return r._beg
}
// End Setter
// 查询结束时间。UTC时间。精度到秒
func (r *TaobaoOpenimAppChatlogsGetRequest) SetEnd(_end int64) error {
    r._end = _end
    r.Set("end", _end)
    return nil
}

// End Getter
func (r TaobaoOpenimAppChatlogsGetRequest) GetEnd() int64 {
    return r._end
}
// Count Setter
// 查询最大条数
func (r *TaobaoOpenimAppChatlogsGetRequest) SetCount(_count int64) error {
    r._count = _count
    r.Set("count", _count)
    return nil
}

// Count Getter
func (r TaobaoOpenimAppChatlogsGetRequest) GetCount() int64 {
    return r._count
}
// Next Setter
// 迭代key
func (r *TaobaoOpenimAppChatlogsGetRequest) SetNext(_next string) error {
    r._next = _next
    r.Set("next", _next)
    return nil
}

// Next Getter
func (r TaobaoOpenimAppChatlogsGetRequest) GetNext() string {
    return r._next
}
