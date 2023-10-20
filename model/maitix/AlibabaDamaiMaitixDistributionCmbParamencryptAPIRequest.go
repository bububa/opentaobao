package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest 加密招商一网能支付入参 API请求
// alibaba.damai.maitix.distribution.cmb.paramencrypt
//
// encryptParam4Cmb
type AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest struct {
	model.Params
	// 入参param
	_param *DisEncrypt4cmbParam
}

// NewAlibabaDamaiMaitixDistributionCmbParamencryptRequest 初始化AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest对象
func NewAlibabaDamaiMaitixDistributionCmbParamencryptRequest() *AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest {
	return &AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.cmb.paramencrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参param
func (r *AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest) SetParam(_param *DisEncrypt4cmbParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest) GetParam() *DisEncrypt4cmbParam {
	return r._param
}
