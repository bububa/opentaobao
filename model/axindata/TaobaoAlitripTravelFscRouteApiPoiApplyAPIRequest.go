package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest 线路供应商提交新增POI申请 API请求
// taobao.alitrip.travel.fsc.route.api.poi.apply
//
// 线路供应商提交新增POI申请
type TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest struct {
	model.Params
	// fscPoiApplyRequest
	_fscPoiApplyRequest *FscPoiApplyRequest
}

// NewTaobaoAlitripTravelFscRouteApiPoiApplyRequest 初始化TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiPoiApplyRequest() *TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest) Reset() {
	r._fscPoiApplyRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.poi.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscPoiApplyRequest is FscPoiApplyRequest Setter
// fscPoiApplyRequest
func (r *TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest) SetFscPoiApplyRequest(_fscPoiApplyRequest *FscPoiApplyRequest) error {
	r._fscPoiApplyRequest = _fscPoiApplyRequest
	r.Set("fsc_poi_apply_request", _fscPoiApplyRequest)
	return nil
}

// GetFscPoiApplyRequest FscPoiApplyRequest Getter
func (r TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest) GetFscPoiApplyRequest() *FscPoiApplyRequest {
	return r._fscPoiApplyRequest
}

var poolTaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiPoiApplyRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiPoiApplyRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest
func GetTaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest() *TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest 将 TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest(v *TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest.Put(v)
}
