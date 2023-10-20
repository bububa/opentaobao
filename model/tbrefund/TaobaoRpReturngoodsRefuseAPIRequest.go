package tbrefund

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpReturngoodsRefuseAPIRequest 卖家拒绝退货 API请求
// taobao.rp.returngoods.refuse
//
// 卖家拒绝退货，目前仅支持天猫退货。
type TaobaoRpReturngoodsRefuseAPIRequest struct {
	model.Params
	// 退款服务状态，售后或者售中
	_refundPhase string
	// 退款编号
	_refundId int64
	// 退款版本号
	_refundVersion int64
	// 拒绝退货凭证图片，必须图片格式，大小不能超过5M
	_refuseProof *model.File
	// 拒绝原因编号，会提供拒绝原因列表供选择
	_refuseReasonId int64
}

// NewTaobaoRpReturngoodsRefuseRequest 初始化TaobaoRpReturngoodsRefuseAPIRequest对象
func NewTaobaoRpReturngoodsRefuseRequest() *TaobaoRpReturngoodsRefuseAPIRequest {
	return &TaobaoRpReturngoodsRefuseAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRpReturngoodsRefuseAPIRequest) Reset() {
	r._refundPhase = ""
	r._refundId = 0
	r._refundVersion = 0
	r._refuseProof = nil
	r._refuseReasonId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetApiMethodName() string {
	return "taobao.rp.returngoods.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundPhase is RefundPhase Setter
// 退款服务状态，售后或者售中
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefundPhase(_refundPhase string) error {
	r._refundPhase = _refundPhase
	r.Set("refund_phase", _refundPhase)
	return nil
}

// GetRefundPhase RefundPhase Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefundPhase() string {
	return r._refundPhase
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefundVersion is RefundVersion Setter
// 退款版本号
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefundVersion(_refundVersion int64) error {
	r._refundVersion = _refundVersion
	r.Set("refund_version", _refundVersion)
	return nil
}

// GetRefundVersion RefundVersion Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefundVersion() int64 {
	return r._refundVersion
}

// SetRefuseProof is RefuseProof Setter
// 拒绝退货凭证图片，必须图片格式，大小不能超过5M
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefuseProof(_refuseProof *model.File) error {
	r._refuseProof = _refuseProof
	r.Set("refuse_proof", _refuseProof)
	return nil
}

// GetRefuseProof RefuseProof Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefuseProof() *model.File {
	return r._refuseProof
}

// SetRefuseReasonId is RefuseReasonId Setter
// 拒绝原因编号，会提供拒绝原因列表供选择
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefuseReasonId(_refuseReasonId int64) error {
	r._refuseReasonId = _refuseReasonId
	r.Set("refuse_reason_id", _refuseReasonId)
	return nil
}

// GetRefuseReasonId RefuseReasonId Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefuseReasonId() int64 {
	return r._refuseReasonId
}

var poolTaobaoRpReturngoodsRefuseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRpReturngoodsRefuseRequest()
	},
}

// GetTaobaoRpReturngoodsRefuseRequest 从 sync.Pool 获取 TaobaoRpReturngoodsRefuseAPIRequest
func GetTaobaoRpReturngoodsRefuseAPIRequest() *TaobaoRpReturngoodsRefuseAPIRequest {
	return poolTaobaoRpReturngoodsRefuseAPIRequest.Get().(*TaobaoRpReturngoodsRefuseAPIRequest)
}

// ReleaseTaobaoRpReturngoodsRefuseAPIRequest 将 TaobaoRpReturngoodsRefuseAPIRequest 放入 sync.Pool
func ReleaseTaobaoRpReturngoodsRefuseAPIRequest(v *TaobaoRpReturngoodsRefuseAPIRequest) {
	v.Reset()
	poolTaobaoRpReturngoodsRefuseAPIRequest.Put(v)
}
