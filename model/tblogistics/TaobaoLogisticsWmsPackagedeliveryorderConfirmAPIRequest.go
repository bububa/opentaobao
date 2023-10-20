package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmspackagedeliveryorderconfirmAPIRequest 包裹出库单确认 API请求
// taobao.logistics.wms.packagedeliveryorder.confirm
//
// 包裹出库单确认
type TaobaologisticswmspackagedeliveryorderconfirmAPIRequest struct {
	model.Params
	// 请求
	_confirmPackageDeliveryOrderRequest *ConfirmPackageOrderRequest
}

// NewTaobaologisticswmspackagedeliveryorderconfirmRequest 初始化TaobaologisticswmspackagedeliveryorderconfirmAPIRequest对象
func NewTaobaologisticswmspackagedeliveryorderconfirmRequest() *TaobaologisticswmspackagedeliveryorderconfirmAPIRequest {
	return &TaobaologisticswmspackagedeliveryorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswmspackagedeliveryorderconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packagedeliveryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswmspackagedeliveryorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswmspackagedeliveryorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmPackageDeliveryOrderRequest is ConfirmPackageDeliveryOrderRequest Setter
// 请求
func (r *TaobaologisticswmspackagedeliveryorderconfirmAPIRequest) SetConfirmPackageDeliveryOrderRequest(_confirmPackageDeliveryOrderRequest *ConfirmPackageOrderRequest) error {
	r._confirmPackageDeliveryOrderRequest = _confirmPackageDeliveryOrderRequest
	r.Set("confirm_package_delivery_order_request", _confirmPackageDeliveryOrderRequest)
	return nil
}

// GetConfirmPackageDeliveryOrderRequest ConfirmPackageDeliveryOrderRequest Getter
func (r TaobaologisticswmspackagedeliveryorderconfirmAPIRequest) GetConfirmPackageDeliveryOrderRequest() *ConfirmPackageOrderRequest {
	return r._confirmPackageDeliveryOrderRequest
}
