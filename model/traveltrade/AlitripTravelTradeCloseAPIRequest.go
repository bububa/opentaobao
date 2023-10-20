package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptraveltradecloseAPIRequest 飞猪度假-订单关闭接口（快速退款） API请求
// alitrip.travel.trade.close
//
// 卖家关单（快速退款接口），不支持二次预约商品的订单
type AlitriptraveltradecloseAPIRequest struct {
	model.Params
	// 交易关闭原因
	_closeReason string
	// 订单关闭原因描述
	_reasonDesc string
	// 子订单编号
	_subOrderId int64
}

// NewAlitriptraveltradecloseRequest 初始化AlitriptraveltradecloseAPIRequest对象
func NewAlitriptraveltradecloseRequest() *AlitriptraveltradecloseAPIRequest {
	return &AlitriptraveltradecloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptraveltradecloseAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptraveltradecloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptraveltradecloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCloseReason is CloseReason Setter
// 交易关闭原因
func (r *AlitriptraveltradecloseAPIRequest) SetCloseReason(_closeReason string) error {
	r._closeReason = _closeReason
	r.Set("close_reason", _closeReason)
	return nil
}

// GetCloseReason CloseReason Getter
func (r AlitriptraveltradecloseAPIRequest) GetCloseReason() string {
	return r._closeReason
}

// SetReasonDesc is ReasonDesc Setter
// 订单关闭原因描述
func (r *AlitriptraveltradecloseAPIRequest) SetReasonDesc(_reasonDesc string) error {
	r._reasonDesc = _reasonDesc
	r.Set("reason_desc", _reasonDesc)
	return nil
}

// GetReasonDesc ReasonDesc Getter
func (r AlitriptraveltradecloseAPIRequest) GetReasonDesc() string {
	return r._reasonDesc
}

// SetSubOrderId is SubOrderId Setter
// 子订单编号
func (r *AlitriptraveltradecloseAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r AlitriptraveltradecloseAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}
