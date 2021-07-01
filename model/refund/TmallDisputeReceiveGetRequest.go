package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫逆向纠纷查询 API请求
tmall.dispute.receive.get

展示商家所有退款信息
*/
type TmallDisputeReceiveGetAPIRequest struct {
    model.Params
    // 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意);WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货);WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货);CLOSED(退款关闭); SUCCESS(退款成功);SELLER_REFUSE_BUYER(卖家拒绝退款);WAIT_BUYER_CONFIRM_REDO_SEND_GOODS(等待买家确认重新邮寄的货物);WAIT_SELLER_CONFIRM_RETURN_ADDRESS(等待卖家确认退货地址);WAIT_SELLER_CONSIGN_GOOGDS(卖家确认收货,等待卖家发货);EXCHANGE_TRANSFORM_TO_REFUND(换货关闭,转退货退款);EXCHANGE_WAIT_BUYER_CONFIRM_GOODS(卖家已发货,等待买家确认收货);POST_FEE_DISPUTE_WAIT_ACTIVATE(邮费单已创建,待激活)
    _status   string
    // 每页条数。取值范围:大于零的整数; 默认值:20;最大值:100
    _pageSize   int64
    // 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
    _useHasNext   bool
    // 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，查看可选值
    _type   string
    // 逆向纠纷单号id
    _refundId   int64
    // 页码。取值范围:大于零的整数; 默认值:1
    _pageNo   int64
    // 买家昵称
    _buyerNick   string
    // 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
    _startModified   string
    // 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
    _endModified   string
    // 需要返回的字段。目前支持有：refund_id, alipay_no, tid, buyer_nick, seller_nick, status, created, modified, order_status, refund_fee, good_status, show_return_logistic(展现买家退货的物流信息), show_exchange_logistic(展现换货的物流信息), time_out, oid, refund_version, title, num, dispute_request, reason, desc
    _fields   []string
}

// 初始化TmallDisputeReceiveGetAPIRequest对象
func NewTmallDisputeReceiveGetRequest() *TmallDisputeReceiveGetAPIRequest{
    return &TmallDisputeReceiveGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDisputeReceiveGetAPIRequest) GetApiMethodName() string {
    return "tmall.dispute.receive.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDisputeReceiveGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意);WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货);WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货);CLOSED(退款关闭); SUCCESS(退款成功);SELLER_REFUSE_BUYER(卖家拒绝退款);WAIT_BUYER_CONFIRM_REDO_SEND_GOODS(等待买家确认重新邮寄的货物);WAIT_SELLER_CONFIRM_RETURN_ADDRESS(等待卖家确认退货地址);WAIT_SELLER_CONSIGN_GOOGDS(卖家确认收货,等待卖家发货);EXCHANGE_TRANSFORM_TO_REFUND(换货关闭,转退货退款);EXCHANGE_WAIT_BUYER_CONFIRM_GOODS(卖家已发货,等待买家确认收货);POST_FEE_DISPUTE_WAIT_ACTIVATE(邮费单已创建,待激活)
func (r *TmallDisputeReceiveGetAPIRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TmallDisputeReceiveGetAPIRequest) GetStatus() string {
    return r._status
}
// PageSize Setter
// 每页条数。取值范围:大于零的整数; 默认值:20;最大值:100
func (r *TmallDisputeReceiveGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallDisputeReceiveGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// UseHasNext Setter
// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
func (r *TmallDisputeReceiveGetAPIRequest) SetUseHasNext(_useHasNext bool) error {
    r._useHasNext = _useHasNext
    r.Set("use_has_next", _useHasNext)
    return nil
}

// UseHasNext Getter
func (r TmallDisputeReceiveGetAPIRequest) GetUseHasNext() bool {
    return r._useHasNext
}
// Type Setter
// 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，查看可选值
func (r *TmallDisputeReceiveGetAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TmallDisputeReceiveGetAPIRequest) GetType() string {
    return r._type
}
// RefundId Setter
// 逆向纠纷单号id
func (r *TmallDisputeReceiveGetAPIRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TmallDisputeReceiveGetAPIRequest) GetRefundId() int64 {
    return r._refundId
}
// PageNo Setter
// 页码。取值范围:大于零的整数; 默认值:1
func (r *TmallDisputeReceiveGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TmallDisputeReceiveGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// BuyerNick Setter
// 买家昵称
func (r *TmallDisputeReceiveGetAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TmallDisputeReceiveGetAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}
// StartModified Setter
// 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
func (r *TmallDisputeReceiveGetAPIRequest) SetStartModified(_startModified string) error {
    r._startModified = _startModified
    r.Set("start_modified", _startModified)
    return nil
}

// StartModified Getter
func (r TmallDisputeReceiveGetAPIRequest) GetStartModified() string {
    return r._startModified
}
// EndModified Setter
// 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
func (r *TmallDisputeReceiveGetAPIRequest) SetEndModified(_endModified string) error {
    r._endModified = _endModified
    r.Set("end_modified", _endModified)
    return nil
}

// EndModified Getter
func (r TmallDisputeReceiveGetAPIRequest) GetEndModified() string {
    return r._endModified
}
// Fields Setter
// 需要返回的字段。目前支持有：refund_id, alipay_no, tid, buyer_nick, seller_nick, status, created, modified, order_status, refund_fee, good_status, show_return_logistic(展现买家退货的物流信息), show_exchange_logistic(展现换货的物流信息), time_out, oid, refund_version, title, num, dispute_request, reason, desc
func (r *TmallDisputeReceiveGetAPIRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TmallDisputeReceiveGetAPIRequest) GetFields() []string {
    return r._fields
}
