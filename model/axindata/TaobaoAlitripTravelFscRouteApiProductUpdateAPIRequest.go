package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest 更新线路产品基本信息 API请求
// taobao.alitrip.travel.fsc.route.api.product.update
//
// 更新线路产品基本信息
type TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest struct {
	model.Params
	// fscRouteProductUpdateRequest
	_fscRouteProductUpdateRequest *FscRouteProductUpdateRequest
}

// NewTaobaoAlitripTravelFscRouteApiProductUpdateRequest 初始化TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProductUpdateRequest() *TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) Reset() {
	r._fscRouteProductUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscRouteProductUpdateRequest is FscRouteProductUpdateRequest Setter
// fscRouteProductUpdateRequest
func (r *TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) SetFscRouteProductUpdateRequest(_fscRouteProductUpdateRequest *FscRouteProductUpdateRequest) error {
	r._fscRouteProductUpdateRequest = _fscRouteProductUpdateRequest
	r.Set("fsc_route_product_update_request", _fscRouteProductUpdateRequest)
	return nil
}

// GetFscRouteProductUpdateRequest FscRouteProductUpdateRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) GetFscRouteProductUpdateRequest() *FscRouteProductUpdateRequest {
	return r._fscRouteProductUpdateRequest
}

var poolTaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiProductUpdateRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductUpdateRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest
func GetTaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest() *TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest 将 TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest(v *TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest.Put(v)
}
