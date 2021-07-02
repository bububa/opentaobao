package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderDistributionCreateAPIRequest 大麦-新分销下单 API请求
// alibaba.damai.maitix.order.distribution.create
//
// createDistributionOrder
type AlibabaDamaiMaitixOrderDistributionCreateAPIRequest struct {
	model.Params
	// 下单参数param
	_param *MoaOrderParam
}

// NewAlibabaDamaiMaitixOrderDistributionCreateRequest 初始化AlibabaDamaiMaitixOrderDistributionCreateAPIRequest对象
func NewAlibabaDamaiMaitixOrderDistributionCreateRequest() *AlibabaDamaiMaitixOrderDistributionCreateAPIRequest {
	return &AlibabaDamaiMaitixOrderDistributionCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.distribution.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 下单参数param
func (r *AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) SetParam(_param *MoaOrderParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) GetParam() *MoaOrderParam {
	return r._param
}
