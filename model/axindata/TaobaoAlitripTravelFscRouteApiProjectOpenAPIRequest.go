package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest 打开团期 API请求
// taobao.alitrip.travel.fsc.route.api.project.open
//
// 打开团期
type TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest struct {
	model.Params
	// fscProjectOpenRequest
	_fscProjectOpenRequest *FscProjectOpenRequest
}

// NewTaobaoAlitripTravelFscRouteApiProjectOpenRequest 初始化TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProjectOpenRequest() *TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest) Reset() {
	r._fscProjectOpenRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.open"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectOpenRequest is FscProjectOpenRequest Setter
// fscProjectOpenRequest
func (r *TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest) SetFscProjectOpenRequest(_fscProjectOpenRequest *FscProjectOpenRequest) error {
	r._fscProjectOpenRequest = _fscProjectOpenRequest
	r.Set("fsc_project_open_request", _fscProjectOpenRequest)
	return nil
}

// GetFscProjectOpenRequest FscProjectOpenRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest) GetFscProjectOpenRequest() *FscProjectOpenRequest {
	return r._fscProjectOpenRequest
}

var poolTaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiProjectOpenRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectOpenRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest
func GetTaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest() *TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest 将 TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest(v *TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest.Put(v)
}
