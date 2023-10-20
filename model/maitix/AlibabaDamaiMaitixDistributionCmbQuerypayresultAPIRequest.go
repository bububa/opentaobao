package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest 查询招行支付状态api API请求
// alibaba.damai.maitix.distribution.cmb.querypayresult
//
// queryPayResult
type AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest struct {
	model.Params
	// 入参param
	_param *QueryPayResultParam
}

// NewAlibabaDamaiMaitixDistributionCmbQuerypayresultRequest 初始化AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest对象
func NewAlibabaDamaiMaitixDistributionCmbQuerypayresultRequest() *AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest {
	return &AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.cmb.querypayresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参param
func (r *AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest) SetParam(_param *QueryPayResultParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest) GetParam() *QueryPayResultParam {
	return r._param
}
