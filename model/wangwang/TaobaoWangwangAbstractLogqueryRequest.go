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
type TaobaoWangwangAbstractLogqueryAPIRequest struct {
    model.Params
    // 买家id，有cntaobao前缀
    _toId   string
    // 卖家id，有cntaobao前缀
    _fromId   string
    // 获取记录条数，默认值是1000
    _count   int64
    // 设置了这个值，那么聊天记录就从这个点开始查询
    _nextKey   string
    // 传入参数的字符集
    _charset   string
    // utc
    _startDate   int64
    // utc
    _endDate   int64
}

// 初始化TaobaoWangwangAbstractLogqueryAPIRequest对象
func NewTaobaoWangwangAbstractLogqueryRequest() *TaobaoWangwangAbstractLogqueryAPIRequest{
    return &TaobaoWangwangAbstractLogqueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWangwangAbstractLogqueryAPIRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.logquery"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWangwangAbstractLogqueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ToId Setter
// 买家id，有cntaobao前缀
func (r *TaobaoWangwangAbstractLogqueryAPIRequest) SetToId(_toId string) error {
    r._toId = _toId
    r.Set("to_id", _toId)
    return nil
}

// ToId Getter
func (r TaobaoWangwangAbstractLogqueryAPIRequest) GetToId() string {
    return r._toId
}
// FromId Setter
// 卖家id，有cntaobao前缀
func (r *TaobaoWangwangAbstractLogqueryAPIRequest) SetFromId(_fromId string) error {
    r._fromId = _fromId
    r.Set("from_id", _fromId)
    return nil
}

// FromId Getter
func (r TaobaoWangwangAbstractLogqueryAPIRequest) GetFromId() string {
    return r._fromId
}
// Count Setter
// 获取记录条数，默认值是1000
func (r *TaobaoWangwangAbstractLogqueryAPIRequest) SetCount(_count int64) error {
    r._count = _count
    r.Set("count", _count)
    return nil
}

// Count Getter
func (r TaobaoWangwangAbstractLogqueryAPIRequest) GetCount() int64 {
    return r._count
}
// NextKey Setter
// 设置了这个值，那么聊天记录就从这个点开始查询
func (r *TaobaoWangwangAbstractLogqueryAPIRequest) SetNextKey(_nextKey string) error {
    r._nextKey = _nextKey
    r.Set("next_key", _nextKey)
    return nil
}

// NextKey Getter
func (r TaobaoWangwangAbstractLogqueryAPIRequest) GetNextKey() string {
    return r._nextKey
}
// Charset Setter
// 传入参数的字符集
func (r *TaobaoWangwangAbstractLogqueryAPIRequest) SetCharset(_charset string) error {
    r._charset = _charset
    r.Set("charset", _charset)
    return nil
}

// Charset Getter
func (r TaobaoWangwangAbstractLogqueryAPIRequest) GetCharset() string {
    return r._charset
}
// StartDate Setter
// utc
func (r *TaobaoWangwangAbstractLogqueryAPIRequest) SetStartDate(_startDate int64) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoWangwangAbstractLogqueryAPIRequest) GetStartDate() int64 {
    return r._startDate
}
// EndDate Setter
// utc
func (r *TaobaoWangwangAbstractLogqueryAPIRequest) SetEndDate(_endDate int64) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoWangwangAbstractLogqueryAPIRequest) GetEndDate() int64 {
    return r._endDate
}
