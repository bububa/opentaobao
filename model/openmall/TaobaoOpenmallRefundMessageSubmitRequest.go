package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交退款单留言 APIRequest
taobao.openmall.refund.message.submit

OpenMall业务提交退款单留言
*/
type TaobaoOpenmallRefundMessageSubmitRequest struct {
    model.Params

    // 分销者身份
    distributor   string 

    // 退款单ID
    refundId   int64 

    // 提交留言结构
    refundMessage   *RefundMessage 

}

func NewTaobaoOpenmallRefundMessageSubmitRequest() *TaobaoOpenmallRefundMessageSubmitRequest{
    return &TaobaoOpenmallRefundMessageSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallRefundMessageSubmitRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.message.submit"
}

func (r TaobaoOpenmallRefundMessageSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallRefundMessageSubmitRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallRefundMessageSubmitRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallRefundMessageSubmitRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoOpenmallRefundMessageSubmitRequest) GetRefundId() int64 {
    return r.refundId
}

func (r *TaobaoOpenmallRefundMessageSubmitRequest) SetRefundMessage(refundMessage *RefundMessage) error {
    r.refundMessage = refundMessage
    r.Set("refund_message", refundMessage)
    return nil
}

func (r TaobaoOpenmallRefundMessageSubmitRequest) GetRefundMessage() *RefundMessage {
    return r.refundMessage
}

