package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionaccountAPIRequest 机票分销企业或者tmc企业预存or月结账户查询接口 API请求
// alitrip.btrip.flight.distribution.account
//
// 机票分销企业或者tmc企业预存or月结账户查询
type AlitripbtripflightdistributionaccountAPIRequest struct {
	model.Params
	// 入参
	_paramAccountRQ *BtripAccountRq
}

// NewAlitripbtripflightdistributionaccountRequest 初始化AlitripbtripflightdistributionaccountAPIRequest对象
func NewAlitripbtripflightdistributionaccountRequest() *AlitripbtripflightdistributionaccountAPIRequest {
	return &AlitripbtripflightdistributionaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionaccountAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAccountRQ is ParamAccountRQ Setter
// 入参
func (r *AlitripbtripflightdistributionaccountAPIRequest) SetParamAccountRQ(_paramAccountRQ *BtripAccountRq) error {
	r._paramAccountRQ = _paramAccountRQ
	r.Set("param_account_r_q", _paramAccountRQ)
	return nil
}

// GetParamAccountRQ ParamAccountRQ Getter
func (r AlitripbtripflightdistributionaccountAPIRequest) GetParamAccountRQ() *BtripAccountRq {
	return r._paramAccountRQ
}
