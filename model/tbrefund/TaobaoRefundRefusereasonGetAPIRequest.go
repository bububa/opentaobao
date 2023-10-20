package tbrefund

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundRefusereasonGetAPIRequest 获取拒绝原因列表 API请求
// taobao.refund.refusereason.get
//
// 获取商家拒绝原因列表
type TaobaoRefundRefusereasonGetAPIRequest struct {
	model.Params
	// 售中或售后
	_refundPhase string
	// 返回参数
	_fields string
	// 退款编号
	_refundId int64
}

// NewTaobaoRefundRefusereasonGetRequest 初始化TaobaoRefundRefusereasonGetAPIRequest对象
func NewTaobaoRefundRefusereasonGetRequest() *TaobaoRefundRefusereasonGetAPIRequest {
	return &TaobaoRefundRefusereasonGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRefundRefusereasonGetAPIRequest) Reset() {
	r._refundPhase = ""
	r._fields = ""
	r._refundId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRefundRefusereasonGetAPIRequest) GetApiMethodName() string {
	return "taobao.refund.refusereason.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRefundRefusereasonGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRefundRefusereasonGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundPhase is RefundPhase Setter
// 售中或售后
func (r *TaobaoRefundRefusereasonGetAPIRequest) SetRefundPhase(_refundPhase string) error {
	r._refundPhase = _refundPhase
	r.Set("refund_phase", _refundPhase)
	return nil
}

// GetRefundPhase RefundPhase Getter
func (r TaobaoRefundRefusereasonGetAPIRequest) GetRefundPhase() string {
	return r._refundPhase
}

// SetFields is Fields Setter
// 返回参数
func (r *TaobaoRefundRefusereasonGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoRefundRefusereasonGetAPIRequest) GetFields() string {
	return r._fields
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaoRefundRefusereasonGetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoRefundRefusereasonGetAPIRequest) GetRefundId() int64 {
	return r._refundId
}

var poolTaobaoRefundRefusereasonGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRefundRefusereasonGetRequest()
	},
}

// GetTaobaoRefundRefusereasonGetRequest 从 sync.Pool 获取 TaobaoRefundRefusereasonGetAPIRequest
func GetTaobaoRefundRefusereasonGetAPIRequest() *TaobaoRefundRefusereasonGetAPIRequest {
	return poolTaobaoRefundRefusereasonGetAPIRequest.Get().(*TaobaoRefundRefusereasonGetAPIRequest)
}

// ReleaseTaobaoRefundRefusereasonGetAPIRequest 将 TaobaoRefundRefusereasonGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoRefundRefusereasonGetAPIRequest(v *TaobaoRefundRefusereasonGetAPIRequest) {
	v.Reset()
	poolTaobaoRefundRefusereasonGetAPIRequest.Put(v)
}
