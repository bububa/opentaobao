package maitix

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.delivery.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixDistributionDeliveryCalculateRequest()
	},
}

// GetAlibabaDamaiMaitixDistributionDeliveryCalculateRequest 从 sync.Pool 获取 AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest
func GetAlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest() *AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest {
	return poolAlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest.Get().(*AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest)
}

// ReleaseAlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest 将 AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest(v *AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest.Put(v)
}
