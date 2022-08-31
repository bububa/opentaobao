package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionAccountAPIRequest 机票分销企业或者tmc企业预存or月结账户查询接口 API请求
// alitrip.btrip.flight.distribution.account
//
// 机票分销企业或者tmc企业预存or月结账户查询
type AlitripBtripFlightDistributionAccountAPIRequest struct {
	model.Params
	// 入参
	_paramAccountRQ *BtripAccountRq
}

// NewAlitripBtripFlightDistributionAccountRequest 初始化AlitripBtripFlightDistributionAccountAPIRequest对象
func NewAlitripBtripFlightDistributionAccountRequest() *AlitripBtripFlightDistributionAccountAPIRequest {
	return &AlitripBtripFlightDistributionAccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionAccountAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionAccountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamAccountRQ is ParamAccountRQ Setter
// 入参
func (r *AlitripBtripFlightDistributionAccountAPIRequest) SetParamAccountRQ(_paramAccountRQ *BtripAccountRq) error {
	r._paramAccountRQ = _paramAccountRQ
	r.Set("param_account_r_q", _paramAccountRQ)
	return nil
}

// GetParamAccountRQ ParamAccountRQ Getter
func (r AlitripBtripFlightDistributionAccountAPIRequest) GetParamAccountRQ() *BtripAccountRq {
	return r._paramAccountRQ
}
