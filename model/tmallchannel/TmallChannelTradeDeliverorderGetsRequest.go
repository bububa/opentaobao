package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发货单列表 API请求
tmall.channel.trade.deliverorder.gets

查询发货单列表
*/
type TmallChannelTradeDeliverorderGetsRequest struct {
    model.Params
    // 发货单单号
    _mainDeliverOrderNo   int64
    // 发货单状态列表
    _orderStatusList   []int64
    // 是否包括子发货单
    _isIncludeSubOrder   bool
    // 每页显示数量
    _pageSize   int64
    // 查询第几页
    _pageNumber   int64
    // 是否分页查询
    _needPagination   bool
    // 渠道
    _channel   int64
}

// 初始化TmallChannelTradeDeliverorderGetsRequest对象
func NewTmallChannelTradeDeliverorderGetsRequest() *TmallChannelTradeDeliverorderGetsRequest{
    return &TmallChannelTradeDeliverorderGetsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeDeliverorderGetsRequest) GetApiMethodName() string {
    return "tmall.channel.trade.deliverorder.gets"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeDeliverorderGetsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainDeliverOrderNo Setter
// 发货单单号
func (r *TmallChannelTradeDeliverorderGetsRequest) SetMainDeliverOrderNo(_mainDeliverOrderNo int64) error {
    r._mainDeliverOrderNo = _mainDeliverOrderNo
    r.Set("main_deliver_order_no", _mainDeliverOrderNo)
    return nil
}

// MainDeliverOrderNo Getter
func (r TmallChannelTradeDeliverorderGetsRequest) GetMainDeliverOrderNo() int64 {
    return r._mainDeliverOrderNo
}
// OrderStatusList Setter
// 发货单状态列表
func (r *TmallChannelTradeDeliverorderGetsRequest) SetOrderStatusList(_orderStatusList []int64) error {
    r._orderStatusList = _orderStatusList
    r.Set("order_status_list", _orderStatusList)
    return nil
}

// OrderStatusList Getter
func (r TmallChannelTradeDeliverorderGetsRequest) GetOrderStatusList() []int64 {
    return r._orderStatusList
}
// IsIncludeSubOrder Setter
// 是否包括子发货单
func (r *TmallChannelTradeDeliverorderGetsRequest) SetIsIncludeSubOrder(_isIncludeSubOrder bool) error {
    r._isIncludeSubOrder = _isIncludeSubOrder
    r.Set("is_include_sub_order", _isIncludeSubOrder)
    return nil
}

// IsIncludeSubOrder Getter
func (r TmallChannelTradeDeliverorderGetsRequest) GetIsIncludeSubOrder() bool {
    return r._isIncludeSubOrder
}
// PageSize Setter
// 每页显示数量
func (r *TmallChannelTradeDeliverorderGetsRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallChannelTradeDeliverorderGetsRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNumber Setter
// 查询第几页
func (r *TmallChannelTradeDeliverorderGetsRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TmallChannelTradeDeliverorderGetsRequest) GetPageNumber() int64 {
    return r._pageNumber
}
// NeedPagination Setter
// 是否分页查询
func (r *TmallChannelTradeDeliverorderGetsRequest) SetNeedPagination(_needPagination bool) error {
    r._needPagination = _needPagination
    r.Set("need_pagination", _needPagination)
    return nil
}

// NeedPagination Getter
func (r TmallChannelTradeDeliverorderGetsRequest) GetNeedPagination() bool {
    return r._needPagination
}
// Channel Setter
// 渠道
func (r *TmallChannelTradeDeliverorderGetsRequest) SetChannel(_channel int64) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TmallChannelTradeDeliverorderGetsRequest) GetChannel() int64 {
    return r._channel
}
