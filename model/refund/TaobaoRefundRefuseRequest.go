package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝退款 API请求
taobao.refund.refuse

卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：<br/>1. 传入的refund_id和相应的tid, oid必须匹配<br/>2. 如果一笔订单只有一笔子订单，则tid必须与oid相同<br/>3. 只有卖家才能执行拒绝退款操作<br/>4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单
*/
type TaobaoRefundRefuseRequest struct {
    model.Params
    // 退款单号
    refundId   int64
    // 拒绝退款时的说明信息，长度2-200
    refuseMessage   string
    // 拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG。天猫退款为必填项。
    refuseProof   []*model.File
    // 可选值为：售中：onsale，售后：aftersale，天猫退款为必填项。
    refundPhase   string
    // 退款版本号，天猫退款为必填项。
    refundVersion   int64
    // 拒绝原因编号，会提供用户拒绝原因列表供选择
    refuseReasonId   int64
}

// 初始化TaobaoRefundRefuseRequest对象
func NewTaobaoRefundRefuseRequest() *TaobaoRefundRefuseRequest{
    return &TaobaoRefundRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRefundRefuseRequest) GetApiMethodName() string {
    return "taobao.refund.refuse"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRefundRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款单号
func (r *TaobaoRefundRefuseRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRefundRefuseRequest) GetRefundId() int64 {
    return r.refundId
}
// RefuseMessage Setter
// 拒绝退款时的说明信息，长度2-200
func (r *TaobaoRefundRefuseRequest) SetRefuseMessage(refuseMessage string) error {
    r.refuseMessage = refuseMessage
    r.Set("refuse_message", refuseMessage)
    return nil
}

// RefuseMessage Getter
func (r TaobaoRefundRefuseRequest) GetRefuseMessage() string {
    return r.refuseMessage
}
// RefuseProof Setter
// 拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG。天猫退款为必填项。
func (r *TaobaoRefundRefuseRequest) SetRefuseProof(refuseProof []*model.File) error {
    r.refuseProof = refuseProof
    r.Set("refuse_proof", refuseProof)
    return nil
}

// RefuseProof Getter
func (r TaobaoRefundRefuseRequest) GetRefuseProof() []*model.File {
    return r.refuseProof
}
// RefundPhase Setter
// 可选值为：售中：onsale，售后：aftersale，天猫退款为必填项。
func (r *TaobaoRefundRefuseRequest) SetRefundPhase(refundPhase string) error {
    r.refundPhase = refundPhase
    r.Set("refund_phase", refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRefundRefuseRequest) GetRefundPhase() string {
    return r.refundPhase
}
// RefundVersion Setter
// 退款版本号，天猫退款为必填项。
func (r *TaobaoRefundRefuseRequest) SetRefundVersion(refundVersion int64) error {
    r.refundVersion = refundVersion
    r.Set("refund_version", refundVersion)
    return nil
}

// RefundVersion Getter
func (r TaobaoRefundRefuseRequest) GetRefundVersion() int64 {
    return r.refundVersion
}
// RefuseReasonId Setter
// 拒绝原因编号，会提供用户拒绝原因列表供选择
func (r *TaobaoRefundRefuseRequest) SetRefuseReasonId(refuseReasonId int64) error {
    r.refuseReasonId = refuseReasonId
    r.Set("refuse_reason_id", refuseReasonId)
    return nil
}

// RefuseReasonId Getter
func (r TaobaoRefundRefuseRequest) GetRefuseReasonId() int64 {
    return r.refuseReasonId
}
