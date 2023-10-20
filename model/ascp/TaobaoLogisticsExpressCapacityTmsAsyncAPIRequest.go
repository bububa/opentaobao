package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresscapacitytmsasyncAPIRequest 上门取退产能信息同步/更新 API请求
// taobao.logistics.express.capacity.tms.async
//
// 上门取退产能信息同步/更新
type TaobaologisticsexpresscapacitytmsasyncAPIRequest struct {
	model.Params
	// 上门取退产能信息同步/更新
	_capacityRequest *CapacityRequest
}

// NewTaobaologisticsexpresscapacitytmsasyncRequest 初始化TaobaologisticsexpresscapacitytmsasyncAPIRequest对象
func NewTaobaologisticsexpresscapacitytmsasyncRequest() *TaobaologisticsexpresscapacitytmsasyncAPIRequest {
	return &TaobaologisticsexpresscapacitytmsasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpresscapacitytmsasyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.capacity.tms.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpresscapacitytmsasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpresscapacitytmsasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCapacityRequest is CapacityRequest Setter
// 上门取退产能信息同步/更新
func (r *TaobaologisticsexpresscapacitytmsasyncAPIRequest) SetCapacityRequest(_capacityRequest *CapacityRequest) error {
	r._capacityRequest = _capacityRequest
	r.Set("capacity_request", _capacityRequest)
	return nil
}

// GetCapacityRequest CapacityRequest Getter
func (r TaobaologisticsexpresscapacitytmsasyncAPIRequest) GetCapacityRequest() *CapacityRequest {
	return r._capacityRequest
}
