package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest 更新团期库存 API请求
// taobao.alitrip.travel.fsc.route.api.project.inventory.update
//
// 更新团期库存
type TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest struct {
	model.Params
	// fscProjectInventoryUpdateRequest
	_fscProjectInventoryUpdateRequest *FscProjectInventoryUpdateRequest
}

// NewTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateRequest 初始化TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateRequest() *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest) Reset() {
	r._fscProjectInventoryUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectInventoryUpdateRequest is FscProjectInventoryUpdateRequest Setter
// fscProjectInventoryUpdateRequest
func (r *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest) SetFscProjectInventoryUpdateRequest(_fscProjectInventoryUpdateRequest *FscProjectInventoryUpdateRequest) error {
	r._fscProjectInventoryUpdateRequest = _fscProjectInventoryUpdateRequest
	r.Set("fsc_project_inventory_update_request", _fscProjectInventoryUpdateRequest)
	return nil
}

// GetFscProjectInventoryUpdateRequest FscProjectInventoryUpdateRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest) GetFscProjectInventoryUpdateRequest() *FscProjectInventoryUpdateRequest {
	return r._fscProjectInventoryUpdateRequest
}

var poolTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest
func GetTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest() *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest 将 TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest(v *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest.Put(v)
}
