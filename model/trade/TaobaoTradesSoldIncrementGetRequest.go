package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家已卖出的增量交易数据（根据修改时间） API请求
taobao.trades.sold.increment.get

搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified <= 1天。
<br/>2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
*/
type TaobaoTradesSoldIncrementGetRequest struct {
    model.Params
    // 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0为处方药未审核状态
    fields   string
    // 查询修改开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss
    startModified   string
    // 查询修改结束时间，必须大于修改开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。
    endModified   string
    // 交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
    status   string
    // 交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。
    type   string
    // 买家昵称
    buyerNick   string
    // 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型
    extType   string
    // 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）
    tag   string
    // 页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span>
    pageNo   int64
    // 每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。
    pageSize   int64
    // 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。
    useHasNext   bool
}

// 初始化TaobaoTradesSoldIncrementGetRequest对象
func NewTaobaoTradesSoldIncrementGetRequest() *TaobaoTradesSoldIncrementGetRequest{
    return &TaobaoTradesSoldIncrementGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradesSoldIncrementGetRequest) GetApiMethodName() string {
    return "taobao.trades.sold.increment.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradesSoldIncrementGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0为处方药未审核状态
func (r *TaobaoTradesSoldIncrementGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetFields() string {
    return r.fields
}
// StartModified Setter
// 查询修改开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoTradesSoldIncrementGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

// StartModified Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetStartModified() string {
    return r.startModified
}
// EndModified Setter
// 查询修改结束时间，必须大于修改开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。
func (r *TaobaoTradesSoldIncrementGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

// EndModified Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetEndModified() string {
    return r.endModified
}
// Status Setter
// 交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
func (r *TaobaoTradesSoldIncrementGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetStatus() string {
    return r.status
}
// Type Setter
// 交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。
func (r *TaobaoTradesSoldIncrementGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetType() string {
    return r.type
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoTradesSoldIncrementGetRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetBuyerNick() string {
    return r.buyerNick
}
// ExtType Setter
// 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型
func (r *TaobaoTradesSoldIncrementGetRequest) SetExtType(extType string) error {
    r.extType = extType
    r.Set("ext_type", extType)
    return nil
}

// ExtType Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetExtType() string {
    return r.extType
}
// Tag Setter
// 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）
func (r *TaobaoTradesSoldIncrementGetRequest) SetTag(tag string) error {
    r.tag = tag
    r.Set("tag", tag)
    return nil
}

// Tag Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetTag() string {
    return r.tag
}
// PageNo Setter
// 页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span>
func (r *TaobaoTradesSoldIncrementGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。
func (r *TaobaoTradesSoldIncrementGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// UseHasNext Setter
// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。
func (r *TaobaoTradesSoldIncrementGetRequest) SetUseHasNext(useHasNext bool) error {
    r.useHasNext = useHasNext
    r.Set("use_has_next", useHasNext)
    return nil
}

// UseHasNext Getter
func (r TaobaoTradesSoldIncrementGetRequest) GetUseHasNext() bool {
    return r.useHasNext
}
