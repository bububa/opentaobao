package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmspackagedeliveryorderpullAPIRequest 包裹出库单拉单 API请求
// taobao.logistics.wms.packagedeliveryorder.pull
//
// 包裹出库单拉单
type TaobaologisticswmspackagedeliveryorderpullAPIRequest struct {
	model.Params
	// 请求
	_pullPackageDeliveryOrderRequest *PullPackageOrderRequest
}

// NewTaobaologisticswmspackagedeliveryorderpullRequest 初始化TaobaologisticswmspackagedeliveryorderpullAPIRequest对象
func NewTaobaologisticswmspackagedeliveryorderpullRequest() *TaobaologisticswmspackagedeliveryorderpullAPIRequest {
	return &TaobaologisticswmspackagedeliveryorderpullAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswmspackagedeliveryorderpullAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packagedeliveryorder.pull"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswmspackagedeliveryorderpullAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswmspackagedeliveryorderpullAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPullPackageDeliveryOrderRequest is PullPackageDeliveryOrderRequest Setter
// 请求
func (r *TaobaologisticswmspackagedeliveryorderpullAPIRequest) SetPullPackageDeliveryOrderRequest(_pullPackageDeliveryOrderRequest *PullPackageOrderRequest) error {
	r._pullPackageDeliveryOrderRequest = _pullPackageDeliveryOrderRequest
	r.Set("pull_package_delivery_order_request", _pullPackageDeliveryOrderRequest)
	return nil
}

// GetPullPackageDeliveryOrderRequest PullPackageDeliveryOrderRequest Getter
func (r TaobaologisticswmspackagedeliveryorderpullAPIRequest) GetPullPackageDeliveryOrderRequest() *PullPackageOrderRequest {
	return r._pullPackageDeliveryOrderRequest
}
