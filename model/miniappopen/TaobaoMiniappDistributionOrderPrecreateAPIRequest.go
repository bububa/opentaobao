package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappdistributionorderprecreateAPIRequest 代商家预创建投放计划 API请求
// taobao.miniapp.distribution.order.precreate
//
// 帮助商家，预创建小程序的投放计划，预创建的投放计划，在商家确认以后，则会生效可用。
type TaobaominiappdistributionorderprecreateAPIRequest struct {
	model.Params
	// 投放计划信息
	_orderRequest *DistributionOrderSaveOpenRequest
}

// NewTaobaominiappdistributionorderprecreateRequest 初始化TaobaominiappdistributionorderprecreateAPIRequest对象
func NewTaobaominiappdistributionorderprecreateRequest() *TaobaominiappdistributionorderprecreateAPIRequest {
	return &TaobaominiappdistributionorderprecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappdistributionorderprecreateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.order.precreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappdistributionorderprecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappdistributionorderprecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRequest is OrderRequest Setter
// 投放计划信息
func (r *TaobaominiappdistributionorderprecreateAPIRequest) SetOrderRequest(_orderRequest *DistributionOrderSaveOpenRequest) error {
	r._orderRequest = _orderRequest
	r.Set("order_request", _orderRequest)
	return nil
}

// GetOrderRequest OrderRequest Getter
func (r TaobaominiappdistributionorderprecreateAPIRequest) GetOrderRequest() *DistributionOrderSaveOpenRequest {
	return r._orderRequest
}
