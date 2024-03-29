package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest 更新团期 API请求
// taobao.alitrip.travel.fsc.route.api.project.update
//
// 更新团期
type TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest struct {
	model.Params
	// fscProjectUpdateRequest
	_fscProjectUpdateRequest *FscProjectModifyRequest
}

// NewTaobaoAlitripTravelFscRouteApiProjectUpdateRequest 初始化TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProjectUpdateRequest() *TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest) Reset() {
	r._fscProjectUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectUpdateRequest is FscProjectUpdateRequest Setter
// fscProjectUpdateRequest
func (r *TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest) SetFscProjectUpdateRequest(_fscProjectUpdateRequest *FscProjectModifyRequest) error {
	r._fscProjectUpdateRequest = _fscProjectUpdateRequest
	r.Set("fsc_project_update_request", _fscProjectUpdateRequest)
	return nil
}

// GetFscProjectUpdateRequest FscProjectUpdateRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest) GetFscProjectUpdateRequest() *FscProjectModifyRequest {
	return r._fscProjectUpdateRequest
}

var poolTaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiProjectUpdateRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectUpdateRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest
func GetTaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest() *TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest 将 TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest(v *TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest.Put(v)
}
