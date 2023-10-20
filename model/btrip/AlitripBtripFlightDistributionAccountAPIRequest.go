package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionAccountAPIRequest) Reset() {
	r._paramAccountRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionAccountAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionAccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionAccountAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripFlightDistributionAccountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionAccountRequest()
	},
}

// GetAlitripBtripFlightDistributionAccountRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionAccountAPIRequest
func GetAlitripBtripFlightDistributionAccountAPIRequest() *AlitripBtripFlightDistributionAccountAPIRequest {
	return poolAlitripBtripFlightDistributionAccountAPIRequest.Get().(*AlitripBtripFlightDistributionAccountAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionAccountAPIRequest 将 AlitripBtripFlightDistributionAccountAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionAccountAPIRequest(v *AlitripBtripFlightDistributionAccountAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionAccountAPIRequest.Put(v)
}
