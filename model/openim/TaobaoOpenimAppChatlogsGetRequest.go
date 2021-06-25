package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim应用聊天记录查询 APIRequest
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

func NewTaobaoOpenimAppChatlogsGetRequest() *TaobaoOpenimAppChatlogsGetRequest{
    return &TaobaoOpenimAppChatlogsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimAppChatlogsGetRequest) GetApiMethodName() string {
    return "taobao.openim.app.chatlogs.get"
}

func (r TaobaoOpenimAppChatlogsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimAppChatlogsGetRequest) SetBeg(beg int64) error {
    r.beg = beg
    r.Set("beg", beg)
    return nil
}

func (r TaobaoOpenimAppChatlogsGetRequest) GetBeg() int64 {
    return r.beg
}

func (r *TaobaoOpenimAppChatlogsGetRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

func (r TaobaoOpenimAppChatlogsGetRequest) GetEnd() int64 {
    return r.end
}

func (r *TaobaoOpenimAppChatlogsGetRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

func (r TaobaoOpenimAppChatlogsGetRequest) GetCount() int64 {
    return r.count
}

func (r *TaobaoOpenimAppChatlogsGetRequest) SetNext(next string) error {
    r.next = next
    r.Set("next", next)
    return nil
}

func (r TaobaoOpenimAppChatlogsGetRequest) GetNext() string {
    return r.next
}

