package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformaccountchargeAPIRequest 资金平台余额账户充值 API请求
// alibaba.fundplatform.account.charge
//
// 资金平台余额账户充值【创建账户&amp;返回付款URL】
type AlibabafundplatformaccountchargeAPIRequest struct {
	model.Params
	// 用户ID
	_paramLong int64
	// 入参对象
	_paramChargeRequest *ChargeRequest
}

// NewAlibabafundplatformaccountchargeRequest 初始化AlibabafundplatformaccountchargeAPIRequest对象
func NewAlibabafundplatformaccountchargeRequest() *AlibabafundplatformaccountchargeAPIRequest {
	return &AlibabafundplatformaccountchargeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformaccountchargeAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.account.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformaccountchargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformaccountchargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLong is ParamLong Setter
// 用户ID
func (r *AlibabafundplatformaccountchargeAPIRequest) SetParamLong(_paramLong int64) error {
	r._paramLong = _paramLong
	r.Set("param_long", _paramLong)
	return nil
}

// GetParamLong ParamLong Getter
func (r AlibabafundplatformaccountchargeAPIRequest) GetParamLong() int64 {
	return r._paramLong
}

// SetParamChargeRequest is ParamChargeRequest Setter
// 入参对象
func (r *AlibabafundplatformaccountchargeAPIRequest) SetParamChargeRequest(_paramChargeRequest *ChargeRequest) error {
	r._paramChargeRequest = _paramChargeRequest
	r.Set("param_charge_request", _paramChargeRequest)
	return nil
}

// GetParamChargeRequest ParamChargeRequest Getter
func (r AlibabafundplatformaccountchargeAPIRequest) GetParamChargeRequest() *ChargeRequest {
	return r._paramChargeRequest
}
