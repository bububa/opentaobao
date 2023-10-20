package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest 包裹出库单拉单 API请求
// taobao.logistics.wms.packagedeliveryorder.pull
//
// 包裹出库单拉单
type TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest struct {
	model.Params
	// 请求
	_pullPackageDeliveryOrderRequest *PullPackageOrderRequest
}

// NewTaobaoLogisticsWmsPackagedeliveryorderPullRequest 初始化TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest对象
func NewTaobaoLogisticsWmsPackagedeliveryorderPullRequest() *TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest {
	return &TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest) Reset() {
	r._pullPackageDeliveryOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packagedeliveryorder.pull"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPullPackageDeliveryOrderRequest is PullPackageDeliveryOrderRequest Setter
// 请求
func (r *TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest) SetPullPackageDeliveryOrderRequest(_pullPackageDeliveryOrderRequest *PullPackageOrderRequest) error {
	r._pullPackageDeliveryOrderRequest = _pullPackageDeliveryOrderRequest
	r.Set("pull_package_delivery_order_request", _pullPackageDeliveryOrderRequest)
	return nil
}

// GetPullPackageDeliveryOrderRequest PullPackageDeliveryOrderRequest Getter
func (r TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest) GetPullPackageDeliveryOrderRequest() *PullPackageOrderRequest {
	return r._pullPackageDeliveryOrderRequest
}

var poolTaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsWmsPackagedeliveryorderPullRequest()
	},
}

// GetTaobaoLogisticsWmsPackagedeliveryorderPullRequest 从 sync.Pool 获取 TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest
func GetTaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest() *TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest {
	return poolTaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest.Get().(*TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest)
}

// ReleaseTaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest 将 TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest(v *TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest.Put(v)
}
