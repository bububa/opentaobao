package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixdistributiondeliveryqueryAPIRequest 查询分销物流单 API请求
// alibaba.damai.maitix.distribution.delivery.query
//
// 渠道查询物流订单
type AlibabadamaimaitixdistributiondeliveryqueryAPIRequest struct {
	model.Params
	// 主订单号
	_mainOrderId string
}

// NewAlibabadamaimaitixdistributiondeliveryqueryRequest 初始化AlibabadamaimaitixdistributiondeliveryqueryAPIRequest对象
func NewAlibabadamaimaitixdistributiondeliveryqueryRequest() *AlibabadamaimaitixdistributiondeliveryqueryAPIRequest {
	return &AlibabadamaimaitixdistributiondeliveryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixdistributiondeliveryqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.delivery.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixdistributiondeliveryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixdistributiondeliveryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 主订单号
func (r *AlibabadamaimaitixdistributiondeliveryqueryAPIRequest) SetMainOrderId(_mainOrderId string) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabadamaimaitixdistributiondeliveryqueryAPIRequest) GetMainOrderId() string {
	return r._mainOrderId
}
