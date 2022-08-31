package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundRefuseAPIRequest 卖家拒绝退款 API请求
// taobao.refund.refuse
//
// 卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：&lt;br/&gt;1. 传入的refund_id和相应的tid, oid必须匹配&lt;br/&gt;2. 如果一笔订单只有一笔子订单，则tid必须与oid相同&lt;br/&gt;3. 只有卖家才能执行拒绝退款操作&lt;br/&gt;4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单
type TaobaoRefundRefuseAPIRequest struct {
	model.Params
	// 拒绝退款时的说明信息，长度2-200
	_refuseMessage string
	// 可选值为：售中：onsale，售后：aftersale，天猫退款为必填项。
	_refundPhase string
	// 退款单号
	_refundId int64
	// 拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG。天猫退款为必填项。
	_refuseProof *model.File
	// 退款版本号，天猫退款为必填项。
	_refundVersion int64
	// 拒绝原因编号，会提供用户拒绝原因列表供选择
	_refuseReasonId int64
}

// NewTaobaoRefundRefuseRequest 初始化TaobaoRefundRefuseAPIRequest对象
func NewTaobaoRefundRefuseRequest() *TaobaoRefundRefuseAPIRequest {
	return &TaobaoRefundRefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRefundRefuseAPIRequest) GetApiMethodName() string {
	return "taobao.refund.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRefundRefuseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefuseMessage is RefuseMessage Setter
// 拒绝退款时的说明信息，长度2-200
func (r *TaobaoRefundRefuseAPIRequest) SetRefuseMessage(_refuseMessage string) error {
	r._refuseMessage = _refuseMessage
	r.Set("refuse_message", _refuseMessage)
	return nil
}

// GetRefuseMessage RefuseMessage Getter
func (r TaobaoRefundRefuseAPIRequest) GetRefuseMessage() string {
	return r._refuseMessage
}

// SetRefundPhase is RefundPhase Setter
// 可选值为：售中：onsale，售后：aftersale，天猫退款为必填项。
func (r *TaobaoRefundRefuseAPIRequest) SetRefundPhase(_refundPhase string) error {
	r._refundPhase = _refundPhase
	r.Set("refund_phase", _refundPhase)
	return nil
}

// GetRefundPhase RefundPhase Getter
func (r TaobaoRefundRefuseAPIRequest) GetRefundPhase() string {
	return r._refundPhase
}

// SetRefundId is RefundId Setter
// 退款单号
func (r *TaobaoRefundRefuseAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoRefundRefuseAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefuseProof is RefuseProof Setter
// 拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG。天猫退款为必填项。
func (r *TaobaoRefundRefuseAPIRequest) SetRefuseProof(_refuseProof *model.File) error {
	r._refuseProof = _refuseProof
	r.Set("refuse_proof", _refuseProof)
	return nil
}

// GetRefuseProof RefuseProof Getter
func (r TaobaoRefundRefuseAPIRequest) GetRefuseProof() *model.File {
	return r._refuseProof
}

// SetRefundVersion is RefundVersion Setter
// 退款版本号，天猫退款为必填项。
func (r *TaobaoRefundRefuseAPIRequest) SetRefundVersion(_refundVersion int64) error {
	r._refundVersion = _refundVersion
	r.Set("refund_version", _refundVersion)
	return nil
}

// GetRefundVersion RefundVersion Getter
func (r TaobaoRefundRefuseAPIRequest) GetRefundVersion() int64 {
	return r._refundVersion
}

// SetRefuseReasonId is RefuseReasonId Setter
// 拒绝原因编号，会提供用户拒绝原因列表供选择
func (r *TaobaoRefundRefuseAPIRequest) SetRefuseReasonId(_refuseReasonId int64) error {
	r._refuseReasonId = _refuseReasonId
	r.Set("refuse_reason_id", _refuseReasonId)
	return nil
}

// GetRefuseReasonId RefuseReasonId Getter
func (r TaobaoRefundRefuseAPIRequest) GetRefuseReasonId() int64 {
	return r._refuseReasonId
}
