package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixorderdistributioncreateAPIRequest 大麦-新分销下单 API请求
// alibaba.damai.maitix.order.distribution.create
//
// createDistributionOrder
type AlibabadamaimaitixorderdistributioncreateAPIRequest struct {
	model.Params
	// 下单参数param
	_param *MoaOrderParam
}

// NewAlibabadamaimaitixorderdistributioncreateRequest 初始化AlibabadamaimaitixorderdistributioncreateAPIRequest对象
func NewAlibabadamaimaitixorderdistributioncreateRequest() *AlibabadamaimaitixorderdistributioncreateAPIRequest {
	return &AlibabadamaimaitixorderdistributioncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixorderdistributioncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.distribution.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixorderdistributioncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixorderdistributioncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 下单参数param
func (r *AlibabadamaimaitixorderdistributioncreateAPIRequest) SetParam(_param *MoaOrderParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabadamaimaitixorderdistributioncreateAPIRequest) GetParam() *MoaOrderParam {
	return r._param
}
