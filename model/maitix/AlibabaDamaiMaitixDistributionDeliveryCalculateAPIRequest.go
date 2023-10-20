package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixdistributiondeliverycalculateAPIRequest 计算渠道用户下单快递费 API请求
// alibaba.damai.maitix.distribution.delivery.calculate
//
// 计算渠道用户下单快递费
type AlibabadamaimaitixdistributiondeliverycalculateAPIRequest struct {
	model.Params
	// 入参
	_param *OpenApiPostFeeParam
}

// NewAlibabadamaimaitixdistributiondeliverycalculateRequest 初始化AlibabadamaimaitixdistributiondeliverycalculateAPIRequest对象
func NewAlibabadamaimaitixdistributiondeliverycalculateRequest() *AlibabadamaimaitixdistributiondeliverycalculateAPIRequest {
	return &AlibabadamaimaitixdistributiondeliverycalculateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixdistributiondeliverycalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.delivery.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixdistributiondeliverycalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixdistributiondeliverycalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabadamaimaitixdistributiondeliverycalculateAPIRequest) SetParam(_param *OpenApiPostFeeParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabadamaimaitixdistributiondeliverycalculateAPIRequest) GetParam() *OpenApiPostFeeParam {
	return r._param
}
