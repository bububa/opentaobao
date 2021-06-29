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
    _isIncludeSubOrder   bool
    // 是否包含主单
    _isIncludeMainOrder   bool
    // 是否包含物流信息
    _isIncludeLogistics   bool
    // 每页显示数量
    _pageSize   int64
    // 查询第几页
    _pageNumber   int64
    // 是否分页查询
    _needPagination   bool
    // 主采购单号
    _mainPurchaseOrderNo   int64
    // 分销商Nick
    _distributorNick   string
    // 渠道编码
    _channel   int64
    // 1-代销；2-经销
    _tradeType   int64
    // 1. 待付款 2.已付款待发货 3.已发货待收货 4.交易完成 5.交易关闭
    _orderStatus   int64
    // 创建时间从
    _createTimeStart   string
    // 创建时间到
    _createTimeEnd   string
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
func (r *TmallChannelTradeOrderGetsRequest) SetIsIncludeSubOrder(_isIncludeSubOrder bool) error {
    r._isIncludeSubOrder = _isIncludeSubOrder
    r.Set("is_include_sub_order", _isIncludeSubOrder)
    return nil
}

// IsIncludeSubOrder Getter
func (r TmallChannelTradeOrderGetsRequest) GetIsIncludeSubOrder() bool {
    return r._isIncludeSubOrder
}
// IsIncludeMainOrder Setter
// 是否包含主单
func (r *TmallChannelTradeOrderGetsRequest) SetIsIncludeMainOrder(_isIncludeMainOrder bool) error {
    r._isIncludeMainOrder = _isIncludeMainOrder
    r.Set("is_include_main_order", _isIncludeMainOrder)
    return nil
}

// IsIncludeMainOrder Getter
func (r TmallChannelTradeOrderGetsRequest) GetIsIncludeMainOrder() bool {
    return r._isIncludeMainOrder
}
// IsIncludeLogistics Setter
// 是否包含物流信息
func (r *TmallChannelTradeOrderGetsRequest) SetIsIncludeLogistics(_isIncludeLogistics bool) error {
    r._isIncludeLogistics = _isIncludeLogistics
    r.Set("is_include_logistics", _isIncludeLogistics)
    return nil
}

// IsIncludeLogistics Getter
func (r TmallChannelTradeOrderGetsRequest) GetIsIncludeLogistics() bool {
    return r._isIncludeLogistics
}
// PageSize Setter
// 每页显示数量
func (r *TmallChannelTradeOrderGetsRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallChannelTradeOrderGetsRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNumber Setter
// 查询第几页
func (r *TmallChannelTradeOrderGetsRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TmallChannelTradeOrderGetsRequest) GetPageNumber() int64 {
    return r._pageNumber
}
// NeedPagination Setter
// 是否分页查询
func (r *TmallChannelTradeOrderGetsRequest) SetNeedPagination(_needPagination bool) error {
    r._needPagination = _needPagination
    r.Set("need_pagination", _needPagination)
    return nil
}

// NeedPagination Getter
func (r TmallChannelTradeOrderGetsRequest) GetNeedPagination() bool {
    return r._needPagination
}
// MainPurchaseOrderNo Setter
// 主采购单号
func (r *TmallChannelTradeOrderGetsRequest) SetMainPurchaseOrderNo(_mainPurchaseOrderNo int64) error {
    r._mainPurchaseOrderNo = _mainPurchaseOrderNo
    r.Set("main_purchase_order_no", _mainPurchaseOrderNo)
    return nil
}

// MainPurchaseOrderNo Getter
func (r TmallChannelTradeOrderGetsRequest) GetMainPurchaseOrderNo() int64 {
    return r._mainPurchaseOrderNo
}
// DistributorNick Setter
// 分销商Nick
func (r *TmallChannelTradeOrderGetsRequest) SetDistributorNick(_distributorNick string) error {
    r._distributorNick = _distributorNick
    r.Set("distributor_nick", _distributorNick)
    return nil
}

// DistributorNick Getter
func (r TmallChannelTradeOrderGetsRequest) GetDistributorNick() string {
    return r._distributorNick
}
// Channel Setter
// 渠道编码
func (r *TmallChannelTradeOrderGetsRequest) SetChannel(_channel int64) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TmallChannelTradeOrderGetsRequest) GetChannel() int64 {
    return r._channel
}
// TradeType Setter
// 1-代销；2-经销
func (r *TmallChannelTradeOrderGetsRequest) SetTradeType(_tradeType int64) error {
    r._tradeType = _tradeType
    r.Set("trade_type", _tradeType)
    return nil
}

// TradeType Getter
func (r TmallChannelTradeOrderGetsRequest) GetTradeType() int64 {
    return r._tradeType
}
// OrderStatus Setter
// 1. 待付款 2.已付款待发货 3.已发货待收货 4.交易完成 5.交易关闭
func (r *TmallChannelTradeOrderGetsRequest) SetOrderStatus(_orderStatus int64) error {
    r._orderStatus = _orderStatus
    r.Set("order_status", _orderStatus)
    return nil
}

// OrderStatus Getter
func (r TmallChannelTradeOrderGetsRequest) GetOrderStatus() int64 {
    return r._orderStatus
}
// CreateTimeStart Setter
// 创建时间从
func (r *TmallChannelTradeOrderGetsRequest) SetCreateTimeStart(_createTimeStart string) error {
    r._createTimeStart = _createTimeStart
    r.Set("create_time_start", _createTimeStart)
    return nil
}

// CreateTimeStart Getter
func (r TmallChannelTradeOrderGetsRequest) GetCreateTimeStart() string {
    return r._createTimeStart
}
// CreateTimeEnd Setter
// 创建时间到
func (r *TmallChannelTradeOrderGetsRequest) SetCreateTimeEnd(_createTimeEnd string) error {
    r._createTimeEnd = _createTimeEnd
    r.Set("create_time_end", _createTimeEnd)
    return nil
}

// CreateTimeEnd Getter
func (r TmallChannelTradeOrderGetsRequest) GetCreateTimeEnd() string {
    return r._createTimeEnd
}
