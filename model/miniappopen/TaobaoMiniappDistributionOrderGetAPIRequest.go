package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappdistributionordergetAPIRequest 小程序投放-查询小程序投放计划信息 API请求
// taobao.miniapp.distribution.order.get
//
// 服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息
type TaobaominiappdistributionordergetAPIRequest struct {
	model.Params
	// 查询入参
	_orderIdRequest *DistributionOrderQueryByIdOpenRequest
}

// NewTaobaominiappdistributionordergetRequest 初始化TaobaominiappdistributionordergetAPIRequest对象
func NewTaobaominiappdistributionordergetRequest() *TaobaominiappdistributionordergetAPIRequest {
	return &TaobaominiappdistributionordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappdistributionordergetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappdistributionordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappdistributionordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderIdRequest is OrderIdRequest Setter
// 查询入参
func (r *TaobaominiappdistributionordergetAPIRequest) SetOrderIdRequest(_orderIdRequest *DistributionOrderQueryByIdOpenRequest) error {
	r._orderIdRequest = _orderIdRequest
	r.Set("order_id_request", _orderIdRequest)
	return nil
}

// GetOrderIdRequest OrderIdRequest Getter
func (r TaobaominiappdistributionordergetAPIRequest) GetOrderIdRequest() *DistributionOrderQueryByIdOpenRequest {
	return r._orderIdRequest
}
