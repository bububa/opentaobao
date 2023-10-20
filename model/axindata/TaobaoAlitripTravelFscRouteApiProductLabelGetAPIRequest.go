package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest 获取线路主题 API请求
// taobao.alitrip.travel.fsc.route.api.product.label.get
//
// 获取线路主题
type TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest struct {
	model.Params
	// fscProductLabelQueryRequest
	_fscProductLabelQueryRequest *FscProductLabelQueryRequest
}

// NewTaobaoAlitripTravelFscRouteApiProductLabelGetRequest 初始化TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProductLabelGetRequest() *TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest) Reset() {
	r._fscProductLabelQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.label.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProductLabelQueryRequest is FscProductLabelQueryRequest Setter
// fscProductLabelQueryRequest
func (r *TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest) SetFscProductLabelQueryRequest(_fscProductLabelQueryRequest *FscProductLabelQueryRequest) error {
	r._fscProductLabelQueryRequest = _fscProductLabelQueryRequest
	r.Set("fsc_product_label_query_request", _fscProductLabelQueryRequest)
	return nil
}

// GetFscProductLabelQueryRequest FscProductLabelQueryRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest) GetFscProductLabelQueryRequest() *FscProductLabelQueryRequest {
	return r._fscProductLabelQueryRequest
}

var poolTaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiProductLabelGetRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductLabelGetRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest
func GetTaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest() *TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest 将 TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest(v *TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest.Put(v)
}
