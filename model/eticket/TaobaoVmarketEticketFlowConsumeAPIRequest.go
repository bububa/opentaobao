package eticket

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVmarketEticketFlowConsumeAPIRequest) Reset() {
	r._outerId = ""
	r._code = ""
	r._operator = ""
	r._bizType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketFlowConsumeAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.flow.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketFlowConsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVmarketEticketFlowConsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoVmarketEticketFlowConsumeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVmarketEticketFlowConsumeRequest()
	},
}

// GetTaobaoVmarketEticketFlowConsumeRequest 从 sync.Pool 获取 TaobaoVmarketEticketFlowConsumeAPIRequest
func GetTaobaoVmarketEticketFlowConsumeAPIRequest() *TaobaoVmarketEticketFlowConsumeAPIRequest {
	return poolTaobaoVmarketEticketFlowConsumeAPIRequest.Get().(*TaobaoVmarketEticketFlowConsumeAPIRequest)
}

// ReleaseTaobaoVmarketEticketFlowConsumeAPIRequest 将 TaobaoVmarketEticketFlowConsumeAPIRequest 放入 sync.Pool
func ReleaseTaobaoVmarketEticketFlowConsumeAPIRequest(v *TaobaoVmarketEticketFlowConsumeAPIRequest) {
	v.Reset()
	poolTaobaoVmarketEticketFlowConsumeAPIRequest.Put(v)
}
