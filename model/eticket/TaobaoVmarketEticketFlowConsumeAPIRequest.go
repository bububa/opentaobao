package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketFlowConsumeAPIRequest 无交易类凭证核销 API请求
// taobao.vmarket.eticket.flow.consume
//
// 无交易类凭证核销
type TaobaoVmarketEticketFlowConsumeAPIRequest struct {
	model.Params
	// 业务单号
	_outerId string
	// 凭证码
	_code string
	// 核销操作人
	_operator string
	// 淘宝业务提供的业务类型值，请联系相关业务运营取得
	_bizType int64
}

// NewTaobaoVmarketEticketFlowConsumeRequest 初始化TaobaoVmarketEticketFlowConsumeAPIRequest对象
func NewTaobaoVmarketEticketFlowConsumeRequest() *TaobaoVmarketEticketFlowConsumeAPIRequest {
	return &TaobaoVmarketEticketFlowConsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketFlowConsumeAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.flow.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketFlowConsumeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterId is OuterId Setter
// 业务单号
func (r *TaobaoVmarketEticketFlowConsumeAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoVmarketEticketFlowConsumeAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetCode is Code Setter
// 凭证码
func (r *TaobaoVmarketEticketFlowConsumeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoVmarketEticketFlowConsumeAPIRequest) GetCode() string {
	return r._code
}

// SetOperator is Operator Setter
// 核销操作人
func (r *TaobaoVmarketEticketFlowConsumeAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoVmarketEticketFlowConsumeAPIRequest) GetOperator() string {
	return r._operator
}

// SetBizType is BizType Setter
// 淘宝业务提供的业务类型值，请联系相关业务运营取得
func (r *TaobaoVmarketEticketFlowConsumeAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoVmarketEticketFlowConsumeAPIRequest) GetBizType() int64 {
	return r._bizType
}
