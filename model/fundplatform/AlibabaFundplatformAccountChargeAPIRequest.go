package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformAccountChargeAPIRequest 资金平台余额账户充值 API请求
// alibaba.fundplatform.account.charge
//
// 资金平台余额账户充值【创建账户&amp;返回付款URL】
type AlibabaFundplatformAccountChargeAPIRequest struct {
	model.Params
	// 用户ID
	_paramLong int64
	// 入参对象
	_paramChargeRequest *ChargeRequest
}

// NewAlibabaFundplatformAccountChargeRequest 初始化AlibabaFundplatformAccountChargeAPIRequest对象
func NewAlibabaFundplatformAccountChargeRequest() *AlibabaFundplatformAccountChargeAPIRequest {
	return &AlibabaFundplatformAccountChargeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformAccountChargeAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.account.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformAccountChargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformAccountChargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLong is ParamLong Setter
// 用户ID
func (r *AlibabaFundplatformAccountChargeAPIRequest) SetParamLong(_paramLong int64) error {
	r._paramLong = _paramLong
	r.Set("param_long", _paramLong)
	return nil
}

// GetParamLong ParamLong Getter
func (r AlibabaFundplatformAccountChargeAPIRequest) GetParamLong() int64 {
	return r._paramLong
}

// SetParamChargeRequest is ParamChargeRequest Setter
// 入参对象
func (r *AlibabaFundplatformAccountChargeAPIRequest) SetParamChargeRequest(_paramChargeRequest *ChargeRequest) error {
	r._paramChargeRequest = _paramChargeRequest
	r.Set("param_charge_request", _paramChargeRequest)
	return nil
}

// GetParamChargeRequest ParamChargeRequest Getter
func (r AlibabaFundplatformAccountChargeAPIRequest) GetParamChargeRequest() *ChargeRequest {
	return r._paramChargeRequest
}
