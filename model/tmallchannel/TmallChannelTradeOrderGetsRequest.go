package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询采购单 API请求
tmall.channel.trade.order.gets

分页查询采购单
*/
type TmallChannelTradeOrderGetsRequest struct {
    model.Params
    // 是否包含子单
    isIncludeSubOrder   bool
    // 是否包含主单
    isIncludeMainOrder   bool
    // 是否包含物流信息
    isIncludeLogistics   bool
    // 每页显示数量
    pageSize   int64
    // 查询第几页
    pageNumber   int64
    // 是否分页查询
    needPagination   bool
    // 主采购单号
    mainPurchaseOrderNo   int64
    // 分销商Nick
    distributorNick   string
    // 渠道编码
    channel   int64
    // 1-代销；2-经销
    tradeType   int64
    // 1. 待付款 2.已付款待发货 3.已发货待收货 4.交易完成 5.交易关闭
    orderStatus   int64
    // 创建时间从
    createTimeStart   string
    // 创建时间到
    createTimeEnd   string
}

// 初始化TmallChannelTradeOrderGetsRequest对象
func NewTmallChannelTradeOrderGetsRequest() *TmallChannelTradeOrderGetsRequest{
    return &TmallChannelTradeOrderGetsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeOrderGetsRequest) GetApiMethodName() string {
    return "tmall.channel.trade.order.gets"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeOrderGetsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsIncludeSubOrder Setter
// 是否包含子单
func (r *TmallChannelTradeOrderGetsRequest) SetIsIncludeSubOrder(isIncludeSubOrder bool) error {
    r.isIncludeSubOrder = isIncludeSubOrder
    r.Set("is_include_sub_order", isIncludeSubOrder)
    return nil
}

// IsIncludeSubOrder Getter
func (r TmallChannelTradeOrderGetsRequest) GetIsIncludeSubOrder() bool {
    return r.isIncludeSubOrder
}
// IsIncludeMainOrder Setter
// 是否包含主单
func (r *TmallChannelTradeOrderGetsRequest) SetIsIncludeMainOrder(isIncludeMainOrder bool) error {
    r.isIncludeMainOrder = isIncludeMainOrder
    r.Set("is_include_main_order", isIncludeMainOrder)
    return nil
}

// IsIncludeMainOrder Getter
func (r TmallChannelTradeOrderGetsRequest) GetIsIncludeMainOrder() bool {
    return r.isIncludeMainOrder
}
// IsIncludeLogistics Setter
// 是否包含物流信息
func (r *TmallChannelTradeOrderGetsRequest) SetIsIncludeLogistics(isIncludeLogistics bool) error {
    r.isIncludeLogistics = isIncludeLogistics
    r.Set("is_include_logistics", isIncludeLogistics)
    return nil
}

// IsIncludeLogistics Getter
func (r TmallChannelTradeOrderGetsRequest) GetIsIncludeLogistics() bool {
    return r.isIncludeLogistics
}
// PageSize Setter
// 每页显示数量
func (r *TmallChannelTradeOrderGetsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TmallChannelTradeOrderGetsRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNumber Setter
// 查询第几页
func (r *TmallChannelTradeOrderGetsRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

// PageNumber Getter
func (r TmallChannelTradeOrderGetsRequest) GetPageNumber() int64 {
    return r.pageNumber
}
// NeedPagination Setter
// 是否分页查询
func (r *TmallChannelTradeOrderGetsRequest) SetNeedPagination(needPagination bool) error {
    r.needPagination = needPagination
    r.Set("need_pagination", needPagination)
    return nil
}

// NeedPagination Getter
func (r TmallChannelTradeOrderGetsRequest) GetNeedPagination() bool {
    return r.needPagination
}
// MainPurchaseOrderNo Setter
// 主采购单号
func (r *TmallChannelTradeOrderGetsRequest) SetMainPurchaseOrderNo(mainPurchaseOrderNo int64) error {
    r.mainPurchaseOrderNo = mainPurchaseOrderNo
    r.Set("main_purchase_order_no", mainPurchaseOrderNo)
    return nil
}

// MainPurchaseOrderNo Getter
func (r TmallChannelTradeOrderGetsRequest) GetMainPurchaseOrderNo() int64 {
    return r.mainPurchaseOrderNo
}
// DistributorNick Setter
// 分销商Nick
func (r *TmallChannelTradeOrderGetsRequest) SetDistributorNick(distributorNick string) error {
    r.distributorNick = distributorNick
    r.Set("distributor_nick", distributorNick)
    return nil
}

// DistributorNick Getter
func (r TmallChannelTradeOrderGetsRequest) GetDistributorNick() string {
    return r.distributorNick
}
// Channel Setter
// 渠道编码
func (r *TmallChannelTradeOrderGetsRequest) SetChannel(channel int64) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r TmallChannelTradeOrderGetsRequest) GetChannel() int64 {
    return r.channel
}
// TradeType Setter
// 1-代销；2-经销
func (r *TmallChannelTradeOrderGetsRequest) SetTradeType(tradeType int64) error {
    r.tradeType = tradeType
    r.Set("trade_type", tradeType)
    return nil
}

// TradeType Getter
func (r TmallChannelTradeOrderGetsRequest) GetTradeType() int64 {
    return r.tradeType
}
// OrderStatus Setter
// 1. 待付款 2.已付款待发货 3.已发货待收货 4.交易完成 5.交易关闭
func (r *TmallChannelTradeOrderGetsRequest) SetOrderStatus(orderStatus int64) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

// OrderStatus Getter
func (r TmallChannelTradeOrderGetsRequest) GetOrderStatus() int64 {
    return r.orderStatus
}
// CreateTimeStart Setter
// 创建时间从
func (r *TmallChannelTradeOrderGetsRequest) SetCreateTimeStart(createTimeStart string) error {
    r.createTimeStart = createTimeStart
    r.Set("create_time_start", createTimeStart)
    return nil
}

// CreateTimeStart Getter
func (r TmallChannelTradeOrderGetsRequest) GetCreateTimeStart() string {
    return r.createTimeStart
}
// CreateTimeEnd Setter
// 创建时间到
func (r *TmallChannelTradeOrderGetsRequest) SetCreateTimeEnd(createTimeEnd string) error {
    r.createTimeEnd = createTimeEnd
    r.Set("create_time_end", createTimeEnd)
    return nil
}

// CreateTimeEnd Getter
func (r TmallChannelTradeOrderGetsRequest) GetCreateTimeEnd() string {
    return r.createTimeEnd
}
