package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest 计算渠道用户下单快递费 API请求
// alibaba.damai.maitix.distribution.delivery.calculate
//
// 计算渠道用户下单快递费
type AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest struct {
	model.Params
	// 入参
	_param *OpenApiPostFeeParam
}

// NewAlibabaDamaiMaitixDistributionDeliveryCalculateRequest 初始化AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest对象
func NewAlibabaDamaiMaitixDistributionDeliveryCalculateRequest() *AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest {
	return &AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.delivery.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest) SetParam(_param *OpenApiPostFeeParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest) GetParam() *OpenApiPostFeeParam {
	return r._param
}
