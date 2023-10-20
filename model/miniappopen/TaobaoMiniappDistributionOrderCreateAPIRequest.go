package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappdistributionordercreateAPIRequest 创建小程序投放计划 API请求
// taobao.miniapp.distribution.order.create
//
// 帮助商家，创建小程序的投放计划。该api，仅针对特定场景开放，目前仅支持客服场景，具体支持的场景列表请联系技术支持或业务对接人确认。
type TaobaominiappdistributionordercreateAPIRequest struct {
	model.Params
	// 投放计划信息
	_orderRequest *DistributionOrderSaveOpenRequest
}

// NewTaobaominiappdistributionordercreateRequest 初始化TaobaominiappdistributionordercreateAPIRequest对象
func NewTaobaominiappdistributionordercreateRequest() *TaobaominiappdistributionordercreateAPIRequest {
	return &TaobaominiappdistributionordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappdistributionordercreateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappdistributionordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappdistributionordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRequest is OrderRequest Setter
// 投放计划信息
func (r *TaobaominiappdistributionordercreateAPIRequest) SetOrderRequest(_orderRequest *DistributionOrderSaveOpenRequest) error {
	r._orderRequest = _orderRequest
	r.Set("order_request", _orderRequest)
	return nil
}

// GetOrderRequest OrderRequest Getter
func (r TaobaominiappdistributionordercreateAPIRequest) GetOrderRequest() *DistributionOrderSaveOpenRequest {
	return r._orderRequest
}
