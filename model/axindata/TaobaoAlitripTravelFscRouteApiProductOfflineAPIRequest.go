package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest 产品下线 API请求
// taobao.alitrip.travel.fsc.route.api.product.offline
//
// 产品下线
type TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest struct {
	model.Params
	// fscProductOfflineRequest
	_fscProductOfflineRequest *FscProductOfflineRequest
}

// NewTaobaoAlitripTravelFscRouteApiProductOfflineRequest 初始化TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProductOfflineRequest() *TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest) Reset() {
	r._fscProductOfflineRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProductOfflineRequest is FscProductOfflineRequest Setter
// fscProductOfflineRequest
func (r *TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest) SetFscProductOfflineRequest(_fscProductOfflineRequest *FscProductOfflineRequest) error {
	r._fscProductOfflineRequest = _fscProductOfflineRequest
	r.Set("fsc_product_offline_request", _fscProductOfflineRequest)
	return nil
}

// GetFscProductOfflineRequest FscProductOfflineRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest) GetFscProductOfflineRequest() *FscProductOfflineRequest {
	return r._fscProductOfflineRequest
}

var poolTaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiProductOfflineRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductOfflineRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest
func GetTaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest() *TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest 将 TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest(v *TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest.Put(v)
}
