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
    beg   int64
    // 查询结束时间。UTC时间。精度到秒
    end   int64
    // 查询最大条数
    count   int64
    // 迭代key
    next   string
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
func (r *TaobaoOpenimAppChatlogsGetRequest) SetBeg(beg int64) error {
    r.beg = beg
    r.Set("beg", beg)
    return nil
}

// Beg Getter
func (r TaobaoOpenimAppChatlogsGetRequest) GetBeg() int64 {
    return r.beg
}
// End Setter
// 查询结束时间。UTC时间。精度到秒
func (r *TaobaoOpenimAppChatlogsGetRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

// End Getter
func (r TaobaoOpenimAppChatlogsGetRequest) GetEnd() int64 {
    return r.end
}
// Count Setter
// 查询最大条数
func (r *TaobaoOpenimAppChatlogsGetRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

// Count Getter
func (r TaobaoOpenimAppChatlogsGetRequest) GetCount() int64 {
    return r.count
}
// Next Setter
// 迭代key
func (r *TaobaoOpenimAppChatlogsGetRequest) SetNext(next string) error {
    r.next = next
    r.Set("next", next)
    return nil
}

// Next Getter
func (r TaobaoOpenimAppChatlogsGetRequest) GetNext() string {
    return r.next
}
