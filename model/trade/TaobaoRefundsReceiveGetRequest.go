package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家收到的退款列表 APIRequest
taobao.refunds.receive.get

查询卖家收到的退款列表
*/
type TaobaoRefundsReceiveGetRequest struct {
    model.Params

    // 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase
    fields   []string 

    // 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)
    status   string 

    // 买家昵称
    buyerNick   string 

    // 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>
    type   string 

    // 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
    startModified   string 

    // 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
    endModified   string 

    // 页码。取值范围:大于零的整数; 默认值:1
    pageNo   int64 

    // 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
    pageSize   int64 

    // 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
    useHasNext   bool 

}

func NewTaobaoRefundsReceiveGetRequest() *TaobaoRefundsReceiveGetRequest{
    return &TaobaoRefundsReceiveGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRefundsReceiveGetRequest) GetApiMethodName() string {
    return "taobao.refunds.receive.get"
}

func (r TaobaoRefundsReceiveGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRefundsReceiveGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoRefundsReceiveGetRequest) GetFields() []string {
    return r.fields
}

func (r *TaobaoRefundsReceiveGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoRefundsReceiveGetRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoRefundsReceiveGetRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoRefundsReceiveGetRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TaobaoRefundsReceiveGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoRefundsReceiveGetRequest) GetType() string {
    return r.type
}

func (r *TaobaoRefundsReceiveGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

func (r TaobaoRefundsReceiveGetRequest) GetStartModified() string {
    return r.startModified
}

func (r *TaobaoRefundsReceiveGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

func (r TaobaoRefundsReceiveGetRequest) GetEndModified() string {
    return r.endModified
}

func (r *TaobaoRefundsReceiveGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoRefundsReceiveGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoRefundsReceiveGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoRefundsReceiveGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoRefundsReceiveGetRequest) SetUseHasNext(useHasNext bool) error {
    r.useHasNext = useHasNext
    r.Set("use_has_next", useHasNext)
    return nil
}

func (r TaobaoRefundsReceiveGetRequest) GetUseHasNext() bool {
    return r.useHasNext
}

