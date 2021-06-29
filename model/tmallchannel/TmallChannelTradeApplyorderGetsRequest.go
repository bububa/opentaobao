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
    tradeType   int64
    // 申请单号
    channelPurchaseApplyOrderNo   string
    // 每页显示数量
    pageSize   int64
    // 查询第几页
    pageNumber   int64
    // 是否需要分页
    needPagination   bool
    // 审核状态列表
    auditStatusList   []int64
    // 分销商nick
    distributorNick   string
    // 渠道
    channel   int64
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
func (r *TmallChannelTradeApplyorderGetsRequest) SetTradeType(tradeType int64) error {
    r.tradeType = tradeType
    r.Set("trade_type", tradeType)
    return nil
}

// TradeType Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetTradeType() int64 {
    return r.tradeType
}
// ChannelPurchaseApplyOrderNo Setter
// 申请单号
func (r *TmallChannelTradeApplyorderGetsRequest) SetChannelPurchaseApplyOrderNo(channelPurchaseApplyOrderNo string) error {
    r.channelPurchaseApplyOrderNo = channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", channelPurchaseApplyOrderNo)
    return nil
}

// ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetChannelPurchaseApplyOrderNo() string {
    return r.channelPurchaseApplyOrderNo
}
// PageSize Setter
// 每页显示数量
func (r *TmallChannelTradeApplyorderGetsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNumber Setter
// 查询第几页
func (r *TmallChannelTradeApplyorderGetsRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

// PageNumber Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetPageNumber() int64 {
    return r.pageNumber
}
// NeedPagination Setter
// 是否需要分页
func (r *TmallChannelTradeApplyorderGetsRequest) SetNeedPagination(needPagination bool) error {
    r.needPagination = needPagination
    r.Set("need_pagination", needPagination)
    return nil
}

// NeedPagination Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetNeedPagination() bool {
    return r.needPagination
}
// AuditStatusList Setter
// 审核状态列表
func (r *TmallChannelTradeApplyorderGetsRequest) SetAuditStatusList(auditStatusList []int64) error {
    r.auditStatusList = auditStatusList
    r.Set("audit_status_list", auditStatusList)
    return nil
}

// AuditStatusList Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetAuditStatusList() []int64 {
    return r.auditStatusList
}
// DistributorNick Setter
// 分销商nick
func (r *TmallChannelTradeApplyorderGetsRequest) SetDistributorNick(distributorNick string) error {
    r.distributorNick = distributorNick
    r.Set("distributor_nick", distributorNick)
    return nil
}

// DistributorNick Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetDistributorNick() string {
    return r.distributorNick
}
// Channel Setter
// 渠道
func (r *TmallChannelTradeApplyorderGetsRequest) SetChannel(channel int64) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r TmallChannelTradeApplyorderGetsRequest) GetChannel() int64 {
    return r.channel
}
