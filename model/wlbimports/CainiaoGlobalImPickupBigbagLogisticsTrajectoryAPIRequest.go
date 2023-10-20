package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest 大包物流轨迹查询 API请求
// cainiao.global.im.pickup.bigbag.logistics.trajectory
//
// 大包物流轨迹查询
type CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest struct {
	model.Params
	// 大包物流轨迹查询请求参数
	_bigbagLogisticsQueryRequest *BigbagLogisticsQueryRequest
}

// NewCainiaoglobalimpickupbigbaglogisticstrajectoryRequest 初始化CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest对象
func NewCainiaoglobalimpickupbigbaglogisticstrajectoryRequest() *CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest {
	return &CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.bigbag.logistics.trajectory"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBigbagLogisticsQueryRequest is BigbagLogisticsQueryRequest Setter
// 大包物流轨迹查询请求参数
func (r *CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest) SetBigbagLogisticsQueryRequest(_bigbagLogisticsQueryRequest *BigbagLogisticsQueryRequest) error {
	r._bigbagLogisticsQueryRequest = _bigbagLogisticsQueryRequest
	r.Set("bigbag_logistics_query_request", _bigbagLogisticsQueryRequest)
	return nil
}

// GetBigbagLogisticsQueryRequest BigbagLogisticsQueryRequest Getter
func (r CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest) GetBigbagLogisticsQueryRequest() *BigbagLogisticsQueryRequest {
	return r._bigbagLogisticsQueryRequest
}
