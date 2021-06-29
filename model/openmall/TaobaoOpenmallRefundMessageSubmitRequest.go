package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交退款单留言 API请求
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

// 初始化TaobaoOpenmallRefundMessageSubmitRequest对象
func NewTaobaoOpenmallRefundMessageSubmitRequest() *TaobaoOpenmallRefundMessageSubmitRequest{
    return &TaobaoOpenmallRefundMessageSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundMessageSubmitRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.message.submit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundMessageSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 分销者身份
func (r *TaobaoOpenmallRefundMessageSubmitRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundMessageSubmitRequest) GetDistributor() string {
    return r.distributor
}
// RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundMessageSubmitRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundMessageSubmitRequest) GetRefundId() int64 {
    return r.refundId
}
// RefundMessage Setter
// 提交留言结构
func (r *TaobaoOpenmallRefundMessageSubmitRequest) SetRefundMessage(refundMessage *RefundMessage) error {
    r.refundMessage = refundMessage
    r.Set("refund_message", refundMessage)
    return nil
}

// RefundMessage Getter
func (r TaobaoOpenmallRefundMessageSubmitRequest) GetRefundMessage() *RefundMessage {
    return r.refundMessage
}
