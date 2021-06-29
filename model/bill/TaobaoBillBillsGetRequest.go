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
    fields   []string
    // 科目编号
    accountId   int64
    // 交易编号
    tradeId   int64
    // 子订单编号
    orderId   int64
    // 开始时间
    startTime   string
    // 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
    endTime   string
    // 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
    timeType   int64
    // 页数,建议不要超过100页,越大性能越低,有可能会超时
    pageNo   int64
    // 每页大小,默认40条,可选范围 ：40~100
    pageSize   int64
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
func (r *TaobaoBillBillsGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoBillBillsGetRequest) GetFields() []string {
    return r.fields
}
// AccountId Setter
// 科目编号
func (r *TaobaoBillBillsGetRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r TaobaoBillBillsGetRequest) GetAccountId() int64 {
    return r.accountId
}
// TradeId Setter
// 交易编号
func (r *TaobaoBillBillsGetRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoBillBillsGetRequest) GetTradeId() int64 {
    return r.tradeId
}
// OrderId Setter
// 子订单编号
func (r *TaobaoBillBillsGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoBillBillsGetRequest) GetOrderId() int64 {
    return r.orderId
}
// StartTime Setter
// 开始时间
func (r *TaobaoBillBillsGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoBillBillsGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
func (r *TaobaoBillBillsGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoBillBillsGetRequest) GetEndTime() string {
    return r.endTime
}
// TimeType Setter
// 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
func (r *TaobaoBillBillsGetRequest) SetTimeType(timeType int64) error {
    r.timeType = timeType
    r.Set("time_type", timeType)
    return nil
}

// TimeType Getter
func (r TaobaoBillBillsGetRequest) GetTimeType() int64 {
    return r.timeType
}
// PageNo Setter
// 页数,建议不要超过100页,越大性能越低,有可能会超时
func (r *TaobaoBillBillsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoBillBillsGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页大小,默认40条,可选范围 ：40~100
func (r *TaobaoBillBillsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBillBillsGetRequest) GetPageSize() int64 {
    return r.pageSize
}
