package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取采购申请单列表 API请求
tmall.channel.trade.applyorder.gets

分页查询采购申请单列表
*/
type TmallChannelTradeApplyorderGetsRequest struct {
    model.Params
    // 交易类型
    _tradeType   int64
    // 申请单号
    _channelPurchaseApplyOrderNo   string
    // 每页显示数量
    _pageSize   int64
    // 查询第几页
    _pageNumber   int64
    // 是否需要分页
    _needPagination   bool
    // 审核状态列表
    _auditStatusList   []int64
    // 分销商nick
    _distributorNick   string
    // 渠道
    _channel   int64
}

// 初始化TmallChannelTradeApplyorderGetsRequest对象
func NewTmallChannelTradeApplyorderGetsRequest() *TmallChannelTradeApplyorderGetsRequest{
    return &TmallChannelTradeApplyorderGetsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeApplyorderGetsRequest) GetApiMethodName() string {
    return "tmall.channel.trade.applyorder.gets"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeApplyorderGetsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeType Setter
// 交易类型
func (r *TmallChannelTradeApplyorderGetsRequest) SetTradeType(_tradeType int64) error {
    r._tradeType = _tradeType
    r.Set("trade_type", _tradeType)
    return nil
}

// TradeType Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetTradeType() int64 {
    return r._tradeType
}
// ChannelPurchaseApplyOrderNo Setter
// 申请单号
func (r *TmallChannelTradeApplyorderGetsRequest) SetChannelPurchaseApplyOrderNo(_channelPurchaseApplyOrderNo string) error {
    r._channelPurchaseApplyOrderNo = _channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", _channelPurchaseApplyOrderNo)
    return nil
}

// ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetChannelPurchaseApplyOrderNo() string {
    return r._channelPurchaseApplyOrderNo
}
// PageSize Setter
// 每页显示数量
func (r *TmallChannelTradeApplyorderGetsRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNumber Setter
// 查询第几页
func (r *TmallChannelTradeApplyorderGetsRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetPageNumber() int64 {
    return r._pageNumber
}
// NeedPagination Setter
// 是否需要分页
func (r *TmallChannelTradeApplyorderGetsRequest) SetNeedPagination(_needPagination bool) error {
    r._needPagination = _needPagination
    r.Set("need_pagination", _needPagination)
    return nil
}

// NeedPagination Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetNeedPagination() bool {
    return r._needPagination
}
// AuditStatusList Setter
// 审核状态列表
func (r *TmallChannelTradeApplyorderGetsRequest) SetAuditStatusList(_auditStatusList []int64) error {
    r._auditStatusList = _auditStatusList
    r.Set("audit_status_list", _auditStatusList)
    return nil
}

// AuditStatusList Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetAuditStatusList() []int64 {
    return r._auditStatusList
}
// DistributorNick Setter
// 分销商nick
func (r *TmallChannelTradeApplyorderGetsRequest) SetDistributorNick(_distributorNick string) error {
    r._distributorNick = _distributorNick
    r.Set("distributor_nick", _distributorNick)
    return nil
}

// DistributorNick Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetDistributorNick() string {
    return r._distributorNick
}
// Channel Setter
// 渠道
func (r *TmallChannelTradeApplyorderGetsRequest) SetChannel(_channel int64) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetChannel() int64 {
    return r._channel
}