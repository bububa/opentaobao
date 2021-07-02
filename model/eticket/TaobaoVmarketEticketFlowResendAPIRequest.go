package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketFlowResendAPIRequest 业务重新触发发码短信 API请求
// taobao.vmarket.eticket.flow.resend
//
// 业务重新触发发码短信
type TaobaoVmarketEticketFlowResendAPIRequest struct {
	model.Params
	// 业务单号
	_outerId string
	// 业务类型值，可联系淘宝业务运营取得具体值
	_bizType int64
}

// NewTaobaoVmarketEticketFlowResendRequest 初始化TaobaoVmarketEticketFlowResendAPIRequest对象
func NewTaobaoVmarketEticketFlowResendRequest() *TaobaoVmarketEticketFlowResendAPIRequest {
	return &TaobaoVmarketEticketFlowResendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketFlowResendAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.flow.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketFlowResendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterId is OuterId Setter
// 业务单号
func (r *TaobaoVmarketEticketFlowResendAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoVmarketEticketFlowResendAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetBizType is BizType Setter
// 业务类型值，可联系淘宝业务运营取得具体值
func (r *TaobaoVmarketEticketFlowResendAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoVmarketEticketFlowResendAPIRequest) GetBizType() int64 {
	return r._bizType
}
