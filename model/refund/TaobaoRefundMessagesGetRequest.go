package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询退款留言/凭证列表 APIRequest
taobao.refund.messages.get

查询退款留言/凭证列表
*/
type TaobaoRefundMessagesGetRequest struct {
    model.Params

    // 需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。
    fields   []string 

    // 退款单号
    refundId   int64 

    // 页码
    pageNo   int64 

    // 每页条数
    pageSize   int64 

    // 退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。
    refundPhase   string 

}

func NewTaobaoRefundMessagesGetRequest() *TaobaoRefundMessagesGetRequest{
    return &TaobaoRefundMessagesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRefundMessagesGetRequest) GetApiMethodName() string {
    return "taobao.refund.messages.get"
}

func (r TaobaoRefundMessagesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRefundMessagesGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoRefundMessagesGetRequest) GetFields() []string {
    return r.fields
}

func (r *TaobaoRefundMessagesGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoRefundMessagesGetRequest) GetRefundId() int64 {
    return r.refundId
}

func (r *TaobaoRefundMessagesGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoRefundMessagesGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoRefundMessagesGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoRefundMessagesGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoRefundMessagesGetRequest) SetRefundPhase(refundPhase string) error {
    r.refundPhase = refundPhase
    r.Set("refund_phase", refundPhase)
    return nil
}

func (r TaobaoRefundMessagesGetRequest) GetRefundPhase() string {
    return r.refundPhase
}

