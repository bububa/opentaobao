package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest 加密招商一网能支付入参 API请求
// alibaba.damai.maitix.distribution.cmb.paramencrypt
//
// encryptParam4Cmb
type AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest struct {
	model.Params
	// 入参param
	_param *DisEncrypt4CmbParam
}

// NewAlibabaDamaiMaitixDistributionCmbParamencryptRequest 初始化AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest对象
func NewAlibabaDamaiMaitixDistributionCmbParamencryptRequest() *AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest {
	return &AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
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
func (r *AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest) SetParam(_param *DisEncrypt4CmbParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest) GetParam() *DisEncrypt4CmbParam {
	return r._param
}

var poolAlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixDistributionCmbParamencryptRequest()
	},
}

// GetAlibabaDamaiMaitixDistributionCmbParamencryptRequest 从 sync.Pool 获取 AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest
func GetAlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest() *AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest {
	return poolAlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest.Get().(*AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest)
}

// ReleaseAlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest 将 AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest(v *AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest.Put(v)
}
