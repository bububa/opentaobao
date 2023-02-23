package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest 查询分销物流单 API请求
// alibaba.damai.maitix.distribution.delivery.query
//
// 渠道查询物流订单
type AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest struct {
	model.Params
	// 主订单号
	_mainOrderId string
}

// NewAlibabaDamaiMaitixDistributionDeliveryQueryRequest 初始化AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest对象
func NewAlibabaDamaiMaitixDistributionDeliveryQueryRequest() *AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest {
	return &AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.delivery.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 主订单号
func (r *AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest) SetMainOrderId(_mainOrderId string) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest) GetMainOrderId() string {
	return r._mainOrderId
}
