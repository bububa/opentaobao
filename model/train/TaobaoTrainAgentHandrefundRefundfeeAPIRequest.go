package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentHandrefundRefundfeeAPIRequest
代理商手动退款接口 API请求
taobao.train.agent.handrefund.refundfee

火车票代理商手动退款接口 */
type TaobaoTrainAgentHandrefundRefundfeeAPIRequest struct {
	model.Params
	// 主订单id
	_mainBizOrderId int64
	// 外部订单号
	_outTradeNo string
	// 退款金额,单位为分
	_refundFee int64
}

// NewTaobaoTrainAgentHandrefundRefundfeeRequest 初始化TaobaoTrainAgentHandrefundRefundfeeAPIRequest对象
func NewTaobaoTrainAgentHandrefundRefundfeeRequest() *TaobaoTrainAgentHandrefundRefundfeeAPIRequest {
	return &TaobaoTrainAgentHandrefundRefundfeeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.handrefund.refundfee"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MainBizOrderId Setter
// 主订单id
func (r *TaobaoTrainAgentHandrefundRefundfeeAPIRequest) SetMainBizOrderId(_mainBizOrderId int64) error {
	r._mainBizOrderId = _mainBizOrderId
	r.Set("main_biz_order_id", _mainBizOrderId)
	return nil
}

// Get MainBizOrderId Getter
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetMainBizOrderId() int64 {
	return r._mainBizOrderId
}

// Set is OutTradeNo Setter
// 外部订单号
func (r *TaobaoTrainAgentHandrefundRefundfeeAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// Get OutTradeNo Getter
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// Set is RefundFee Setter
// 退款金额,单位为分
func (r *TaobaoTrainAgentHandrefundRefundfeeAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// Get RefundFee Getter
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}
