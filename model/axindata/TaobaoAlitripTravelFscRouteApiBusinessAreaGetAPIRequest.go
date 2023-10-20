package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest 获取业务区域 API请求
// taobao.alitrip.travel.fsc.route.api.business.area.get
//
// 获取业务区域
type TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest struct {
	model.Params
	// fscBusinessAreaQueryRequest
	_fscBusinessAreaQueryRequest *FscBusinessAreaQueryRequest
}

// NewTaobaoAlitripTravelFscRouteApiBusinessAreaGetRequest 初始化TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiBusinessAreaGetRequest() *TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest) Reset() {
	r._fscBusinessAreaQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.business.area.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscBusinessAreaQueryRequest is FscBusinessAreaQueryRequest Setter
// fscBusinessAreaQueryRequest
func (r *TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest) SetFscBusinessAreaQueryRequest(_fscBusinessAreaQueryRequest *FscBusinessAreaQueryRequest) error {
	r._fscBusinessAreaQueryRequest = _fscBusinessAreaQueryRequest
	r.Set("fsc_business_area_query_request", _fscBusinessAreaQueryRequest)
	return nil
}

// GetFscBusinessAreaQueryRequest FscBusinessAreaQueryRequest Getter
func (r TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest) GetFscBusinessAreaQueryRequest() *FscBusinessAreaQueryRequest {
	return r._fscBusinessAreaQueryRequest
}

var poolTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiBusinessAreaGetRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiBusinessAreaGetRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest
func GetTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest() *TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest 将 TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest(v *TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest.Put(v)
}
