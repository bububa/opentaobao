package refund

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRefundRefusereasonGetAPIRequest) GetApiMethodName() string {
	return "taobao.refund.refusereason.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRefundRefusereasonGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
