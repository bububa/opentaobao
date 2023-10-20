package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest 大包物流轨迹查询 API请求
// cainiao.global.im.pickup.bigbag.logistics.trajectory
//
// 大包物流轨迹查询
type CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest struct {
	model.Params
	// 大包物流轨迹查询请求参数
	_bigbagLogisticsQueryRequest *BigbagLogisticsQueryRequest
}

// NewCainiaoGlobalImPickupBigbagLogisticsTrajectoryRequest 初始化CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest对象
func NewCainiaoGlobalImPickupBigbagLogisticsTrajectoryRequest() *CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest {
	return &CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.bigbag.logistics.trajectory"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBigbagLogisticsQueryRequest is BigbagLogisticsQueryRequest Setter
// 大包物流轨迹查询请求参数
func (r *CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest) SetBigbagLogisticsQueryRequest(_bigbagLogisticsQueryRequest *BigbagLogisticsQueryRequest) error {
	r._bigbagLogisticsQueryRequest = _bigbagLogisticsQueryRequest
	r.Set("bigbag_logistics_query_request", _bigbagLogisticsQueryRequest)
	return nil
}

// GetBigbagLogisticsQueryRequest BigbagLogisticsQueryRequest Getter
func (r CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest) GetBigbagLogisticsQueryRequest() *BigbagLogisticsQueryRequest {
	return r._bigbagLogisticsQueryRequest
}
