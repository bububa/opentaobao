package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询采购退款 API请求
taobao.fenxiao.refund.query

供应商按查询条件批量查询代销采购退款
*/
type TaobaoFenxiaoRefundQueryRequest struct {
    model.Params
    // 代销采购退款单最早修改时间
    startDate   string
    // 代销采购退款最迟修改时间。与start_date的最大时间间隔不能超过30天
    endDate   string
    // 页码（大于0的整数。无值或小于1的值按默认值1计）
    pageNo   int64
    // 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
    pageSize   int64
    // 是否查询下游买家的退款信息
    querySellerRefund   bool
    // 渠道code，可批量 老供销渠道：999
    tradeTypes   []int64
    // 角色，供应商：2，分销商：1
    userRoleType   int64
    // 代销：1 经销：2 寄售（猫超自营寄售）：5 平台寄售：6
    channelCodes   []int64
}

// 初始化TaobaoFenxiaoRefundQueryRequest对象
func NewTaobaoFenxiaoRefundQueryRequest() *TaobaoFenxiaoRefundQueryRequest{
    return &TaobaoFenxiaoRefundQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoRefundQueryRequest) GetApiMethodName() string {
    return "taobao.fenxiao.refund.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoRefundQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 代销采购退款单最早修改时间
func (r *TaobaoFenxiaoRefundQueryRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoFenxiaoRefundQueryRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 代销采购退款最迟修改时间。与start_date的最大时间间隔不能超过30天
func (r *TaobaoFenxiaoRefundQueryRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoFenxiaoRefundQueryRequest) GetEndDate() string {
    return r.endDate
}
// PageNo Setter
// 页码（大于0的整数。无值或小于1的值按默认值1计）
func (r *TaobaoFenxiaoRefundQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoRefundQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
func (r *TaobaoFenxiaoRefundQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoRefundQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// QuerySellerRefund Setter
// 是否查询下游买家的退款信息
func (r *TaobaoFenxiaoRefundQueryRequest) SetQuerySellerRefund(querySellerRefund bool) error {
    r.querySellerRefund = querySellerRefund
    r.Set("query_seller_refund", querySellerRefund)
    return nil
}

// QuerySellerRefund Getter
func (r TaobaoFenxiaoRefundQueryRequest) GetQuerySellerRefund() bool {
    return r.querySellerRefund
}
// TradeTypes Setter
// 渠道code，可批量 老供销渠道：999
func (r *TaobaoFenxiaoRefundQueryRequest) SetTradeTypes(tradeTypes []int64) error {
    r.tradeTypes = tradeTypes
    r.Set("trade_types", tradeTypes)
    return nil
}

// TradeTypes Getter
func (r TaobaoFenxiaoRefundQueryRequest) GetTradeTypes() []int64 {
    return r.tradeTypes
}
// UserRoleType Setter
// 角色，供应商：2，分销商：1
func (r *TaobaoFenxiaoRefundQueryRequest) SetUserRoleType(userRoleType int64) error {
    r.userRoleType = userRoleType
    r.Set("user_role_type", userRoleType)
    return nil
}

// UserRoleType Getter
func (r TaobaoFenxiaoRefundQueryRequest) GetUserRoleType() int64 {
    return r.userRoleType
}
// ChannelCodes Setter
// 代销：1 经销：2 寄售（猫超自营寄售）：5 平台寄售：6
func (r *TaobaoFenxiaoRefundQueryRequest) SetChannelCodes(channelCodes []int64) error {
    r.channelCodes = channelCodes
    r.Set("channel_codes", channelCodes)
    return nil
}

// ChannelCodes Getter
func (r TaobaoFenxiaoRefundQueryRequest) GetChannelCodes() []int64 {
    return r.channelCodes
}
