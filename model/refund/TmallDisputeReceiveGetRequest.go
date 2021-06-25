package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫逆向纠纷查询 APIRequest
tmall.dispute.receive.get

展示商家所有退款信息
*/
type TmallDisputeReceiveGetRequest struct {
    model.Params

    // 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意);WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货);WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货);CLOSED(退款关闭); SUCCESS(退款成功);SELLER_REFUSE_BUYER(卖家拒绝退款);WAIT_BUYER_CONFIRM_REDO_SEND_GOODS(等待买家确认重新邮寄的货物);WAIT_SELLER_CONFIRM_RETURN_ADDRESS(等待卖家确认退货地址);WAIT_SELLER_CONSIGN_GOOGDS(卖家确认收货,等待卖家发货);EXCHANGE_TRANSFORM_TO_REFUND(换货关闭,转退货退款);EXCHANGE_WAIT_BUYER_CONFIRM_GOODS(卖家已发货,等待买家确认收货);POST_FEE_DISPUTE_WAIT_ACTIVATE(邮费单已创建,待激活)
    status   string 

    // 每页条数。取值范围:大于零的整数; 默认值:20;最大值:100
    pageSize   int64 

    // 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
    useHasNext   bool 

    // 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，查看可选值
    type   string 

    // 逆向纠纷单号id
    refundId   int64 

    // 页码。取值范围:大于零的整数; 默认值:1
    pageNo   int64 

    // 买家昵称
    buyerNick   string 

    // 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
    startModified   string 

    // 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
    endModified   string 

    // 需要返回的字段。目前支持有：refund_id, alipay_no, tid, buyer_nick, seller_nick, status, created, modified, order_status, refund_fee, good_status, show_return_logistic(展现买家退货的物流信息), show_exchange_logistic(展现换货的物流信息), time_out, oid, refund_version, title, num, dispute_request, reason, desc
    fields   []String 

}

func NewTmallDisputeReceiveGetRequest() *TmallDisputeReceiveGetRequest{
    return &TmallDisputeReceiveGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallDisputeReceiveGetRequest) GetApiMethodName() string {
    return "tmall.dispute.receive.get"
}

func (r TmallDisputeReceiveGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallDisputeReceiveGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetStatus() string {
    return r.status
}

func (r *TmallDisputeReceiveGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TmallDisputeReceiveGetRequest) SetUseHasNext(useHasNext bool) error {
    r.useHasNext = useHasNext
    r.Set("use_has_next", useHasNext)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetUseHasNext() bool {
    return r.useHasNext
}

func (r *TmallDisputeReceiveGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetType() string {
    return r.type
}

func (r *TmallDisputeReceiveGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetRefundId() int64 {
    return r.refundId
}

func (r *TmallDisputeReceiveGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TmallDisputeReceiveGetRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TmallDisputeReceiveGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetStartModified() string {
    return r.startModified
}

func (r *TmallDisputeReceiveGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetEndModified() string {
    return r.endModified
}

func (r *TmallDisputeReceiveGetRequest) SetFields(fields []String) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TmallDisputeReceiveGetRequest) GetFields() []String {
    return r.fields
}

