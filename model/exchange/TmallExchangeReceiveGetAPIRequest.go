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
type TmallExchangeReceiveGetAPIRequest struct {
    model.Params
    // 查询修改时间段的结束时间点
    _endGmtModifedTime   string
    // 查询修改时间段的开始时间点
    _startGmtModifiedTime   string
    // 快递单号
    _logisticNo   string
    // 买家昵称
    _buyerNick   string
    // 查询申请时间段的开始时间点
    _startCreatedTime   string
    // 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
    _fields   []string
    // 每页条数
    _pageSize   int64
    // 换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14)
    _disputeStatusArray   []int64
    // 查询申请时间段的结束时间点
    _endCreatedTime   string
    // 买家id
    _buyerId   int64
    // 退款单号ID列表，最多只能输入20个id
    _refundIdArray   []int64
    // 页码
    _pageNo   int64
    // 正向订单号
    _bizOrderId   int64
}

// 初始化TmallExchangeReceiveGetAPIRequest对象
func NewTmallExchangeReceiveGetRequest() *TmallExchangeReceiveGetAPIRequest{
    return &TmallExchangeReceiveGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeReceiveGetAPIRequest) GetApiMethodName() string {
    return "tmall.exchange.receive.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeReceiveGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EndGmtModifedTime Setter
// 查询修改时间段的结束时间点
func (r *TmallExchangeReceiveGetAPIRequest) SetEndGmtModifedTime(_endGmtModifedTime string) error {
    r._endGmtModifedTime = _endGmtModifedTime
    r.Set("end_gmt_modifed_time", _endGmtModifedTime)
    return nil
}

// EndGmtModifedTime Getter
func (r TmallExchangeReceiveGetAPIRequest) GetEndGmtModifedTime() string {
    return r._endGmtModifedTime
}
// StartGmtModifiedTime Setter
// 查询修改时间段的开始时间点
func (r *TmallExchangeReceiveGetAPIRequest) SetStartGmtModifiedTime(_startGmtModifiedTime string) error {
    r._startGmtModifiedTime = _startGmtModifiedTime
    r.Set("start_gmt_modified_time", _startGmtModifiedTime)
    return nil
}

// StartGmtModifiedTime Getter
func (r TmallExchangeReceiveGetAPIRequest) GetStartGmtModifiedTime() string {
    return r._startGmtModifiedTime
}
// LogisticNo Setter
// 快递单号
func (r *TmallExchangeReceiveGetAPIRequest) SetLogisticNo(_logisticNo string) error {
    r._logisticNo = _logisticNo
    r.Set("logistic_no", _logisticNo)
    return nil
}

// LogisticNo Getter
func (r TmallExchangeReceiveGetAPIRequest) GetLogisticNo() string {
    return r._logisticNo
}
// BuyerNick Setter
// 买家昵称
func (r *TmallExchangeReceiveGetAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TmallExchangeReceiveGetAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}
// StartCreatedTime Setter
// 查询申请时间段的开始时间点
func (r *TmallExchangeReceiveGetAPIRequest) SetStartCreatedTime(_startCreatedTime string) error {
    r._startCreatedTime = _startCreatedTime
    r.Set("start_created_time", _startCreatedTime)
    return nil
}

// StartCreatedTime Getter
func (r TmallExchangeReceiveGetAPIRequest) GetStartCreatedTime() string {
    return r._startCreatedTime
}
// Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
func (r *TmallExchangeReceiveGetAPIRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TmallExchangeReceiveGetAPIRequest) GetFields() []string {
    return r._fields
}
// PageSize Setter
// 每页条数
func (r *TmallExchangeReceiveGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallExchangeReceiveGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// DisputeStatusArray Setter
// 换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14)
func (r *TmallExchangeReceiveGetAPIRequest) SetDisputeStatusArray(_disputeStatusArray []int64) error {
    r._disputeStatusArray = _disputeStatusArray
    r.Set("dispute_status_array", _disputeStatusArray)
    return nil
}

// DisputeStatusArray Getter
func (r TmallExchangeReceiveGetAPIRequest) GetDisputeStatusArray() []int64 {
    return r._disputeStatusArray
}
// EndCreatedTime Setter
// 查询申请时间段的结束时间点
func (r *TmallExchangeReceiveGetAPIRequest) SetEndCreatedTime(_endCreatedTime string) error {
    r._endCreatedTime = _endCreatedTime
    r.Set("end_created_time", _endCreatedTime)
    return nil
}

// EndCreatedTime Getter
func (r TmallExchangeReceiveGetAPIRequest) GetEndCreatedTime() string {
    return r._endCreatedTime
}
// BuyerId Setter
// 买家id
func (r *TmallExchangeReceiveGetAPIRequest) SetBuyerId(_buyerId int64) error {
    r._buyerId = _buyerId
    r.Set("buyer_id", _buyerId)
    return nil
}

// BuyerId Getter
func (r TmallExchangeReceiveGetAPIRequest) GetBuyerId() int64 {
    return r._buyerId
}
// RefundIdArray Setter
// 退款单号ID列表，最多只能输入20个id
func (r *TmallExchangeReceiveGetAPIRequest) SetRefundIdArray(_refundIdArray []int64) error {
    r._refundIdArray = _refundIdArray
    r.Set("refund_id_array", _refundIdArray)
    return nil
}

// RefundIdArray Getter
func (r TmallExchangeReceiveGetAPIRequest) GetRefundIdArray() []int64 {
    return r._refundIdArray
}
// PageNo Setter
// 页码
func (r *TmallExchangeReceiveGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TmallExchangeReceiveGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// BizOrderId Setter
// 正向订单号
func (r *TmallExchangeReceiveGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TmallExchangeReceiveGetAPIRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
