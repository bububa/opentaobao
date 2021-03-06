package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradeCloseAPIRequest 飞猪度假-订单关闭接口（快速退款） API请求
// alitrip.travel.trade.close
//
// 卖家关单（快速退款接口），不支持二次预约商品的订单
type AlitripTravelTradeCloseAPIRequest struct {
	model.Params
	// 交易关闭原因
	_closeReason string
	// 订单关闭原因描述
	_reasonDesc string
	// 子订单编号
	_subOrderId int64
}

// NewAlitripTravelTradeCloseRequest 初始化AlitripTravelTradeCloseAPIRequest对象
func NewAlitripTravelTradeCloseRequest() *AlitripTravelTradeCloseAPIRequest {
	return &AlitripTravelTradeCloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeCloseAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeCloseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCloseReason is CloseReason Setter
// 交易关闭原因
func (r *AlitripTravelTradeCloseAPIRequest) SetCloseReason(_closeReason string) error {
	r._closeReason = _closeReason
	r.Set("close_reason", _closeReason)
	return nil
}

// GetCloseReason CloseReason Getter
func (r AlitripTravelTradeCloseAPIRequest) GetCloseReason() string {
	return r._closeReason
}

// SetReasonDesc is ReasonDesc Setter
// 订单关闭原因描述
func (r *AlitripTravelTradeCloseAPIRequest) SetReasonDesc(_reasonDesc string) error {
	r._reasonDesc = _reasonDesc
	r.Set("reason_desc", _reasonDesc)
	return nil
}

// GetReasonDesc ReasonDesc Getter
func (r AlitripTravelTradeCloseAPIRequest) GetReasonDesc() string {
	return r._reasonDesc
}

// SetSubOrderId is SubOrderId Setter
// 子订单编号
func (r *AlitripTravelTradeCloseAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r AlitripTravelTradeCloseAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}
