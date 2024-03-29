package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest 产品上线 API请求
// taobao.alitrip.travel.fsc.route.api.product.online
//
// 产品上线
type TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest struct {
	model.Params
	// fscProductOnlineRequest
	_fscProductOnlineRequest *FscProductOnlineRequest
}

// NewTaobaoAlitripTravelFscRouteApiProductOnlineRequest 初始化TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProductOnlineRequest() *TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest) Reset() {
	r._fscProductOnlineRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.online"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProductOnlineRequest is FscProductOnlineRequest Setter
// fscProductOnlineRequest
func (r *TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest) SetFscProductOnlineRequest(_fscProductOnlineRequest *FscProductOnlineRequest) error {
	r._fscProductOnlineRequest = _fscProductOnlineRequest
	r.Set("fsc_product_online_request", _fscProductOnlineRequest)
	return nil
}

// GetFscProductOnlineRequest FscProductOnlineRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest) GetFscProductOnlineRequest() *FscProductOnlineRequest {
	return r._fscProductOnlineRequest
}

var poolTaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiProductOnlineRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductOnlineRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest
func GetTaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest() *TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest 将 TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest(v *TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest.Put(v)
}
