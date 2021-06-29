package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询换货列表 API请求
tmall.exchange.receive.get

卖家查询换货列表
*/
type TmallExchangeReceiveGetRequest struct {
    model.Params
    // 查询修改时间段的结束时间点
    endGmtModifedTime   string
    // 查询修改时间段的开始时间点
    startGmtModifiedTime   string
    // 快递单号
    logisticNo   string
    // 买家昵称
    buyerNick   string
    // 查询申请时间段的开始时间点
    startCreatedTime   string
    // 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
    fields   []string
    // 每页条数
    pageSize   int64
    // 换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14)
    disputeStatusArray   []int64
    // 查询申请时间段的结束时间点
    endCreatedTime   string
    // 买家id
    buyerId   int64
    // 退款单号ID列表，最多只能输入20个id
    refundIdArray   []int64
    // 页码
    pageNo   int64
    // 正向订单号
    bizOrderId   int64
}

// 初始化TmallExchangeReceiveGetRequest对象
func NewTmallExchangeReceiveGetRequest() *TmallExchangeReceiveGetRequest{
    return &TmallExchangeReceiveGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeReceiveGetRequest) GetApiMethodName() string {
    return "tmall.exchange.receive.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeReceiveGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EndGmtModifedTime Setter
// 查询修改时间段的结束时间点
func (r *TmallExchangeReceiveGetRequest) SetEndGmtModifedTime(endGmtModifedTime string) error {
    r.endGmtModifedTime = endGmtModifedTime
    r.Set("end_gmt_modifed_time", endGmtModifedTime)
    return nil
}

// EndGmtModifedTime Getter
func (r TmallExchangeReceiveGetRequest) GetEndGmtModifedTime() string {
    return r.endGmtModifedTime
}
// StartGmtModifiedTime Setter
// 查询修改时间段的开始时间点
func (r *TmallExchangeReceiveGetRequest) SetStartGmtModifiedTime(startGmtModifiedTime string) error {
    r.startGmtModifiedTime = startGmtModifiedTime
    r.Set("start_gmt_modified_time", startGmtModifiedTime)
    return nil
}

// StartGmtModifiedTime Getter
func (r TmallExchangeReceiveGetRequest) GetStartGmtModifiedTime() string {
    return r.startGmtModifiedTime
}
// LogisticNo Setter
// 快递单号
func (r *TmallExchangeReceiveGetRequest) SetLogisticNo(logisticNo string) error {
    r.logisticNo = logisticNo
    r.Set("logistic_no", logisticNo)
    return nil
}

// LogisticNo Getter
func (r TmallExchangeReceiveGetRequest) GetLogisticNo() string {
    return r.logisticNo
}
// BuyerNick Setter
// 买家昵称
func (r *TmallExchangeReceiveGetRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TmallExchangeReceiveGetRequest) GetBuyerNick() string {
    return r.buyerNick
}
// StartCreatedTime Setter
// 查询申请时间段的开始时间点
func (r *TmallExchangeReceiveGetRequest) SetStartCreatedTime(startCreatedTime string) error {
    r.startCreatedTime = startCreatedTime
    r.Set("start_created_time", startCreatedTime)
    return nil
}

// StartCreatedTime Getter
func (r TmallExchangeReceiveGetRequest) GetStartCreatedTime() string {
    return r.startCreatedTime
}
// Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
func (r *TmallExchangeReceiveGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TmallExchangeReceiveGetRequest) GetFields() []string {
    return r.fields
}
// PageSize Setter
// 每页条数
func (r *TmallExchangeReceiveGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TmallExchangeReceiveGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// DisputeStatusArray Setter
// 换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14)
func (r *TmallExchangeReceiveGetRequest) SetDisputeStatusArray(disputeStatusArray []int64) error {
    r.disputeStatusArray = disputeStatusArray
    r.Set("dispute_status_array", disputeStatusArray)
    return nil
}

// DisputeStatusArray Getter
func (r TmallExchangeReceiveGetRequest) GetDisputeStatusArray() []int64 {
    return r.disputeStatusArray
}
// EndCreatedTime Setter
// 查询申请时间段的结束时间点
func (r *TmallExchangeReceiveGetRequest) SetEndCreatedTime(endCreatedTime string) error {
    r.endCreatedTime = endCreatedTime
    r.Set("end_created_time", endCreatedTime)
    return nil
}

// EndCreatedTime Getter
func (r TmallExchangeReceiveGetRequest) GetEndCreatedTime() string {
    return r.endCreatedTime
}
// BuyerId Setter
// 买家id
func (r *TmallExchangeReceiveGetRequest) SetBuyerId(buyerId int64) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

// BuyerId Getter
func (r TmallExchangeReceiveGetRequest) GetBuyerId() int64 {
    return r.buyerId
}
// RefundIdArray Setter
// 退款单号ID列表，最多只能输入20个id
func (r *TmallExchangeReceiveGetRequest) SetRefundIdArray(refundIdArray []int64) error {
    r.refundIdArray = refundIdArray
    r.Set("refund_id_array", refundIdArray)
    return nil
}

// RefundIdArray Getter
func (r TmallExchangeReceiveGetRequest) GetRefundIdArray() []int64 {
    return r.refundIdArray
}
// PageNo Setter
// 页码
func (r *TmallExchangeReceiveGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TmallExchangeReceiveGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// BizOrderId Setter
// 正向订单号
func (r *TmallExchangeReceiveGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TmallExchangeReceiveGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
