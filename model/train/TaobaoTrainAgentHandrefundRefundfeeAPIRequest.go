package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentHandrefundRefundfeeAPIRequest 代理商手动退款接口 API请求
// taobao.train.agent.handrefund.refundfee
//
// 火车票代理商手动退款接口
type TaobaoTrainAgentHandrefundRefundfeeAPIRequest struct {
	model.Params
	// 外部订单号
	_outTradeNo string
	// 主订单id
	_mainBizOrderId int64
	// 退款金额,单位为分
	_refundFee int64
	// 子订单号
	_subOrderId int64
}

// NewTaobaoTrainAgentHandrefundRefundfeeRequest 初始化TaobaoTrainAgentHandrefundRefundfeeAPIRequest对象
func NewTaobaoTrainAgentHandrefundRefundfeeRequest() *TaobaoTrainAgentHandrefundRefundfeeAPIRequest {
	return &TaobaoTrainAgentHandrefundRefundfeeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentHandrefundRefundfeeAPIRequest) Reset() {
	r._outTradeNo = ""
	r._mainBizOrderId = 0
	r._refundFee = 0
	r._subOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.handrefund.refundfee"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutTradeNo is OutTradeNo Setter
// 外部订单号
func (r *TaobaoTrainAgentHandrefundRefundfeeAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetMainBizOrderId is MainBizOrderId Setter
// 主订单id
func (r *TaobaoTrainAgentHandrefundRefundfeeAPIRequest) SetMainBizOrderId(_mainBizOrderId int64) error {
	r._mainBizOrderId = _mainBizOrderId
	r.Set("main_biz_order_id", _mainBizOrderId)
	return nil
}

// GetMainBizOrderId MainBizOrderId Getter
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetMainBizOrderId() int64 {
	return r._mainBizOrderId
}

// SetRefundFee is RefundFee Setter
// 退款金额,单位为分
func (r *TaobaoTrainAgentHandrefundRefundfeeAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// GetRefundFee RefundFee Getter
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}

// SetSubOrderId is SubOrderId Setter
// 子订单号
func (r *TaobaoTrainAgentHandrefundRefundfeeAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r TaobaoTrainAgentHandrefundRefundfeeAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

var poolTaobaoTrainAgentHandrefundRefundfeeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentHandrefundRefundfeeRequest()
	},
}

// GetTaobaoTrainAgentHandrefundRefundfeeRequest 从 sync.Pool 获取 TaobaoTrainAgentHandrefundRefundfeeAPIRequest
func GetTaobaoTrainAgentHandrefundRefundfeeAPIRequest() *TaobaoTrainAgentHandrefundRefundfeeAPIRequest {
	return poolTaobaoTrainAgentHandrefundRefundfeeAPIRequest.Get().(*TaobaoTrainAgentHandrefundRefundfeeAPIRequest)
}

// ReleaseTaobaoTrainAgentHandrefundRefundfeeAPIRequest 将 TaobaoTrainAgentHandrefundRefundfeeAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentHandrefundRefundfeeAPIRequest(v *TaobaoTrainAgentHandrefundRefundfeeAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentHandrefundRefundfeeAPIRequest.Put(v)
}
