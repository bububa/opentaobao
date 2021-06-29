package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取采购申请单列表 APIRequest
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

func NewTmallChannelTradeApplyorderGetsRequest() *TmallChannelTradeApplyorderGetsRequest{
    return &TmallChannelTradeApplyorderGetsRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeApplyorderGetsRequest) GetApiMethodName() string {
    return "tmall.channel.trade.applyorder.gets"
}

func (r TmallChannelTradeApplyorderGetsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeApplyorderGetsRequest) SetTradeType(tradeType int64) error {
    r.tradeType = tradeType
    r.Set("trade_type", tradeType)
    return nil
}

func (r TmallChannelTradeApplyorderGetsRequest) GetTradeType() int64 {
    return r.tradeType
}

func (r *TmallChannelTradeApplyorderGetsRequest) SetChannelPurchaseApplyOrderNo(channelPurchaseApplyOrderNo string) error {
    r.channelPurchaseApplyOrderNo = channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", channelPurchaseApplyOrderNo)
    return nil
}

func (r TmallChannelTradeApplyorderGetsRequest) GetChannelPurchaseApplyOrderNo() string {
    return r.channelPurchaseApplyOrderNo
}

func (r *TmallChannelTradeApplyorderGetsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TmallChannelTradeApplyorderGetsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TmallChannelTradeApplyorderGetsRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

func (r TmallChannelTradeApplyorderGetsRequest) GetPageNumber() int64 {
    return r.pageNumber
}

func (r *TmallChannelTradeApplyorderGetsRequest) SetNeedPagination(needPagination bool) error {
    r.needPagination = needPagination
    r.Set("need_pagination", needPagination)
    return nil
}

func (r TmallChannelTradeApplyorderGetsRequest) GetNeedPagination() bool {
    return r.needPagination
}

func (r *TmallChannelTradeApplyorderGetsRequest) SetAuditStatusList(auditStatusList []int64) error {
    r.auditStatusList = auditStatusList
    r.Set("audit_status_list", auditStatusList)
    return nil
}

func (r TmallChannelTradeApplyorderGetsRequest) GetAuditStatusList() []int64 {
    return r.auditStatusList
}

func (r *TmallChannelTradeApplyorderGetsRequest) SetDistributorNick(distributorNick string) error {
    r.distributorNick = distributorNick
    r.Set("distributor_nick", distributorNick)
    return nil
}

func (r TmallChannelTradeApplyorderGetsRequest) GetDistributorNick() string {
    return r.distributorNick
}

func (r *TmallChannelTradeApplyorderGetsRequest) SetChannel(channel int64) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TmallChannelTradeApplyorderGetsRequest) GetChannel() int64 {
    return r.channel
}

