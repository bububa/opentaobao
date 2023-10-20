package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundrefusereasongetAPIRequest 获取拒绝原因列表 API请求
// taobao.refund.refusereason.get
//
// 获取商家拒绝原因列表
type TaobaorefundrefusereasongetAPIRequest struct {
	model.Params
	// 售中或售后
	_refundPhase string
	// 返回参数
	_fields string
	// 退款编号
	_refundId int64
}

// NewTaobaorefundrefusereasongetRequest 初始化TaobaorefundrefusereasongetAPIRequest对象
func NewTaobaorefundrefusereasongetRequest() *TaobaorefundrefusereasongetAPIRequest {
	return &TaobaorefundrefusereasongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorefundrefusereasongetAPIRequest) GetApiMethodName() string {
	return "taobao.refund.refusereason.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorefundrefusereasongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorefundrefusereasongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundPhase is RefundPhase Setter
// 售中或售后
func (r *TaobaorefundrefusereasongetAPIRequest) SetRefundPhase(_refundPhase string) error {
	r._refundPhase = _refundPhase
	r.Set("refund_phase", _refundPhase)
	return nil
}

// GetRefundPhase RefundPhase Getter
func (r TaobaorefundrefusereasongetAPIRequest) GetRefundPhase() string {
	return r._refundPhase
}

// SetFields is Fields Setter
// 返回参数
func (r *TaobaorefundrefusereasongetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaorefundrefusereasongetAPIRequest) GetFields() string {
	return r._fields
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaorefundrefusereasongetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaorefundrefusereasongetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
