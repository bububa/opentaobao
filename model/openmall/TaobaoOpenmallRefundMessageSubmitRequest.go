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
    _distributor   string
    // 退款单ID
    _refundId   int64
    // 提交留言结构
    _refundMessage   *RefundMessage
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
func (r *TaobaoOpenmallRefundMessageSubmitRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundMessageSubmitRequest) GetDistributor() string {
    return r._distributor
}
// RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundMessageSubmitRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundMessageSubmitRequest) GetRefundId() int64 {
    return r._refundId
}
// RefundMessage Setter
// 提交留言结构
func (r *TaobaoOpenmallRefundMessageSubmitRequest) SetRefundMessage(_refundMessage *RefundMessage) error {
    r._refundMessage = _refundMessage
    r.Set("refund_message", _refundMessage)
    return nil
}

// RefundMessage Getter
func (r TaobaoOpenmallRefundMessageSubmitRequest) GetRefundMessage() *RefundMessage {
    return r._refundMessage
}
