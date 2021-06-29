package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账单明细数据(自研发商家专用) API请求
taobao.bill.bills.get

查询账单明细数据
*/
type TaobaoBillBillsGetRequest struct {
    model.Params
    // 传入需要返回的字段,参见Bill结构体
    _fields   []string
    // 科目编号
    _accountId   int64
    // 交易编号
    _tradeId   int64
    // 子订单编号
    _orderId   int64
    // 开始时间
    _startTime   string
    // 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
    _endTime   string
    // 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
    _timeType   int64
    // 页数,建议不要超过100页,越大性能越低,有可能会超时
    _pageNo   int64
    // 每页大小,默认40条,可选范围 ：40~100
    _pageSize   int64
}

// 初始化TaobaoBillBillsGetRequest对象
func NewTaobaoBillBillsGetRequest() *TaobaoBillBillsGetRequest{
    return &TaobaoBillBillsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBillBillsGetRequest) GetApiMethodName() string {
    return "taobao.bill.bills.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBillBillsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 传入需要返回的字段,参见Bill结构体
func (r *TaobaoBillBillsGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoBillBillsGetRequest) GetFields() []string {
    return r._fields
}
// AccountId Setter
// 科目编号
func (r *TaobaoBillBillsGetRequest) SetAccountId(_accountId int64) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r TaobaoBillBillsGetRequest) GetAccountId() int64 {
    return r._accountId
}
// TradeId Setter
// 交易编号
func (r *TaobaoBillBillsGetRequest) SetTradeId(_tradeId int64) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoBillBillsGetRequest) GetTradeId() int64 {
    return r._tradeId
}
// OrderId Setter
// 子订单编号
func (r *TaobaoBillBillsGetRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoBillBillsGetRequest) GetOrderId() int64 {
    return r._orderId
}
// StartTime Setter
// 开始时间
func (r *TaobaoBillBillsGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoBillBillsGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
func (r *TaobaoBillBillsGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoBillBillsGetRequest) GetEndTime() string {
    return r._endTime
}
// TimeType Setter
// 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
func (r *TaobaoBillBillsGetRequest) SetTimeType(_timeType int64) error {
    r._timeType = _timeType
    r.Set("time_type", _timeType)
    return nil
}

// TimeType Getter
func (r TaobaoBillBillsGetRequest) GetTimeType() int64 {
    return r._timeType
}
// PageNo Setter
// 页数,建议不要超过100页,越大性能越低,有可能会超时
func (r *TaobaoBillBillsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoBillBillsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小,默认40条,可选范围 ：40~100
func (r *TaobaoBillBillsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBillBillsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
