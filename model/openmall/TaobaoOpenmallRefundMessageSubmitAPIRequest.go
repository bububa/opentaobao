package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundMessageSubmitAPIRequest 提交退款单留言 API请求
// taobao.openmall.refund.message.submit
//
// OpenMall业务提交退款单留言
type TaobaoOpenmallRefundMessageSubmitAPIRequest struct {
	model.Params
	// 分销者身份
	_distributor string
	// 退款单ID
	_refundId int64
	// 提交留言结构
	_refundMessage *RefundMessage
}

// NewTaobaoOpenmallRefundMessageSubmitRequest 初始化TaobaoOpenmallRefundMessageSubmitAPIRequest对象
func NewTaobaoOpenmallRefundMessageSubmitRequest() *TaobaoOpenmallRefundMessageSubmitAPIRequest {
	return &TaobaoOpenmallRefundMessageSubmitAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallRefundMessageSubmitAPIRequest) Reset() {
	r._distributor = ""
	r._refundId = 0
	r._refundMessage = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundMessageSubmitAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.message.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundMessageSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallRefundMessageSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 分销者身份
func (r *TaobaoOpenmallRefundMessageSubmitAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallRefundMessageSubmitAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundMessageSubmitAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoOpenmallRefundMessageSubmitAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefundMessage is RefundMessage Setter
// 提交留言结构
func (r *TaobaoOpenmallRefundMessageSubmitAPIRequest) SetRefundMessage(_refundMessage *RefundMessage) error {
	r._refundMessage = _refundMessage
	r.Set("refund_message", _refundMessage)
	return nil
}

// GetRefundMessage RefundMessage Getter
func (r TaobaoOpenmallRefundMessageSubmitAPIRequest) GetRefundMessage() *RefundMessage {
	return r._refundMessage
}

var poolTaobaoOpenmallRefundMessageSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallRefundMessageSubmitRequest()
	},
}

// GetTaobaoOpenmallRefundMessageSubmitRequest 从 sync.Pool 获取 TaobaoOpenmallRefundMessageSubmitAPIRequest
func GetTaobaoOpenmallRefundMessageSubmitAPIRequest() *TaobaoOpenmallRefundMessageSubmitAPIRequest {
	return poolTaobaoOpenmallRefundMessageSubmitAPIRequest.Get().(*TaobaoOpenmallRefundMessageSubmitAPIRequest)
}

// ReleaseTaobaoOpenmallRefundMessageSubmitAPIRequest 将 TaobaoOpenmallRefundMessageSubmitAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallRefundMessageSubmitAPIRequest(v *TaobaoOpenmallRefundMessageSubmitAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallRefundMessageSubmitAPIRequest.Put(v)
}
