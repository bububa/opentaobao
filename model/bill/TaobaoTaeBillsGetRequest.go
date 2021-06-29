package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询账单明细 API请求
taobao.tae.bills.get

tae查询账单明细
*/
type TaobaoTaeBillsGetRequest struct {
    model.Params
    // 开始时间
    queryStartDate   string
    // 交易编号
    pTradeId   int64
    // 子订单编号
    tradeId   int64
    // 每页大小,默认40条,可选范围 ：40~100
    pageSize   int64
    // 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
    queryDateType   int64
    // 页数,建议不要超过100页,越大性能越低,有可能会超时
    currentPage   int64
    // 科目编号
    itemId   int64
    // 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
    queryEndDate   string
    // 传入需要返回的字段,参见Bill结构体
    fields   []string
}

// 初始化TaobaoTaeBillsGetRequest对象
func NewTaobaoTaeBillsGetRequest() *TaobaoTaeBillsGetRequest{
    return &TaobaoTaeBillsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaeBillsGetRequest) GetApiMethodName() string {
    return "taobao.tae.bills.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaeBillsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryStartDate Setter
// 开始时间
func (r *TaobaoTaeBillsGetRequest) SetQueryStartDate(queryStartDate string) error {
    r.queryStartDate = queryStartDate
    r.Set("query_start_date", queryStartDate)
    return nil
}

// QueryStartDate Getter
func (r TaobaoTaeBillsGetRequest) GetQueryStartDate() string {
    return r.queryStartDate
}
// PTradeId Setter
// 交易编号
func (r *TaobaoTaeBillsGetRequest) SetPTradeId(pTradeId int64) error {
    r.pTradeId = pTradeId
    r.Set("p_trade_id", pTradeId)
    return nil
}

// PTradeId Getter
func (r TaobaoTaeBillsGetRequest) GetPTradeId() int64 {
    return r.pTradeId
}
// TradeId Setter
// 子订单编号
func (r *TaobaoTaeBillsGetRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoTaeBillsGetRequest) GetTradeId() int64 {
    return r.tradeId
}
// PageSize Setter
// 每页大小,默认40条,可选范围 ：40~100
func (r *TaobaoTaeBillsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTaeBillsGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// QueryDateType Setter
// 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
func (r *TaobaoTaeBillsGetRequest) SetQueryDateType(queryDateType int64) error {
    r.queryDateType = queryDateType
    r.Set("query_date_type", queryDateType)
    return nil
}

// QueryDateType Getter
func (r TaobaoTaeBillsGetRequest) GetQueryDateType() int64 {
    return r.queryDateType
}
// CurrentPage Setter
// 页数,建议不要超过100页,越大性能越低,有可能会超时
func (r *TaobaoTaeBillsGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoTaeBillsGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// ItemId Setter
// 科目编号
func (r *TaobaoTaeBillsGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoTaeBillsGetRequest) GetItemId() int64 {
    return r.itemId
}
// QueryEndDate Setter
// 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
func (r *TaobaoTaeBillsGetRequest) SetQueryEndDate(queryEndDate string) error {
    r.queryEndDate = queryEndDate
    r.Set("query_end_date", queryEndDate)
    return nil
}

// QueryEndDate Getter
func (r TaobaoTaeBillsGetRequest) GetQueryEndDate() string {
    return r.queryEndDate
}
// Fields Setter
// 传入需要返回的字段,参见Bill结构体
func (r *TaobaoTaeBillsGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoTaeBillsGetRequest) GetFields() []string {
    return r.fields
}
