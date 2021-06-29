package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
模糊聊天记录查询 API请求
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

// 初始化TaobaoWangwangAbstractLogqueryRequest对象
func NewTaobaoWangwangAbstractLogqueryRequest() *TaobaoWangwangAbstractLogqueryRequest{
    return &TaobaoWangwangAbstractLogqueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWangwangAbstractLogqueryRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.logquery"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWangwangAbstractLogqueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ToId Setter
// 买家id，有cntaobao前缀
func (r *TaobaoWangwangAbstractLogqueryRequest) SetToId(toId string) error {
    r.toId = toId
    r.Set("to_id", toId)
    return nil
}

// ToId Getter
func (r TaobaoWangwangAbstractLogqueryRequest) GetToId() string {
    return r.toId
}
// FromId Setter
// 卖家id，有cntaobao前缀
func (r *TaobaoWangwangAbstractLogqueryRequest) SetFromId(fromId string) error {
    r.fromId = fromId
    r.Set("from_id", fromId)
    return nil
}

// FromId Getter
func (r TaobaoWangwangAbstractLogqueryRequest) GetFromId() string {
    return r.fromId
}
// Count Setter
// 获取记录条数，默认值是1000
func (r *TaobaoWangwangAbstractLogqueryRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

// Count Getter
func (r TaobaoWangwangAbstractLogqueryRequest) GetCount() int64 {
    return r.count
}
// NextKey Setter
// 设置了这个值，那么聊天记录就从这个点开始查询
func (r *TaobaoWangwangAbstractLogqueryRequest) SetNextKey(nextKey string) error {
    r.nextKey = nextKey
    r.Set("next_key", nextKey)
    return nil
}

// NextKey Getter
func (r TaobaoWangwangAbstractLogqueryRequest) GetNextKey() string {
    return r.nextKey
}
// Charset Setter
// 传入参数的字符集
func (r *TaobaoWangwangAbstractLogqueryRequest) SetCharset(charset string) error {
    r.charset = charset
    r.Set("charset", charset)
    return nil
}

// Charset Getter
func (r TaobaoWangwangAbstractLogqueryRequest) GetCharset() string {
    return r.charset
}
// StartDate Setter
// utc
func (r *TaobaoWangwangAbstractLogqueryRequest) SetStartDate(startDate int64) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoWangwangAbstractLogqueryRequest) GetStartDate() int64 {
    return r.startDate
}
// EndDate Setter
// utc
func (r *TaobaoWangwangAbstractLogqueryRequest) SetEndDate(endDate int64) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoWangwangAbstractLogqueryRequest) GetEndDate() int64 {
    return r.endDate
}
