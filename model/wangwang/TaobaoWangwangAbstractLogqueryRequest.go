package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
模糊聊天记录查询 APIRequest
taobao.wangwang.abstract.logquery

模糊聊天记录查询
*/
type TaobaoWangwangAbstractLogqueryRequest struct {
    model.Params

    // 买家id，有cntaobao前缀
    toId   string 

    // 卖家id，有cntaobao前缀
    fromId   string 

    // 获取记录条数，默认值是1000
    count   int64 

    // 设置了这个值，那么聊天记录就从这个点开始查询
    nextKey   string 

    // 传入参数的字符集
    charset   string 

    // utc
    startDate   int64 

    // utc
    endDate   int64 

}

func NewTaobaoWangwangAbstractLogqueryRequest() *TaobaoWangwangAbstractLogqueryRequest{
    return &TaobaoWangwangAbstractLogqueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWangwangAbstractLogqueryRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.logquery"
}

func (r TaobaoWangwangAbstractLogqueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWangwangAbstractLogqueryRequest) SetToId(toId string) error {
    r.toId = toId
    r.Set("to_id", toId)
    return nil
}

func (r TaobaoWangwangAbstractLogqueryRequest) GetToId() string {
    return r.toId
}

func (r *TaobaoWangwangAbstractLogqueryRequest) SetFromId(fromId string) error {
    r.fromId = fromId
    r.Set("from_id", fromId)
    return nil
}

func (r TaobaoWangwangAbstractLogqueryRequest) GetFromId() string {
    return r.fromId
}

func (r *TaobaoWangwangAbstractLogqueryRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

func (r TaobaoWangwangAbstractLogqueryRequest) GetCount() int64 {
    return r.count
}

func (r *TaobaoWangwangAbstractLogqueryRequest) SetNextKey(nextKey string) error {
    r.nextKey = nextKey
    r.Set("next_key", nextKey)
    return nil
}

func (r TaobaoWangwangAbstractLogqueryRequest) GetNextKey() string {
    return r.nextKey
}

func (r *TaobaoWangwangAbstractLogqueryRequest) SetCharset(charset string) error {
    r.charset = charset
    r.Set("charset", charset)
    return nil
}

func (r TaobaoWangwangAbstractLogqueryRequest) GetCharset() string {
    return r.charset
}

func (r *TaobaoWangwangAbstractLogqueryRequest) SetStartDate(startDate int64) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoWangwangAbstractLogqueryRequest) GetStartDate() int64 {
    return r.startDate
}

func (r *TaobaoWangwangAbstractLogqueryRequest) SetEndDate(endDate int64) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoWangwangAbstractLogqueryRequest) GetEndDate() int64 {
    return r.endDate
}

