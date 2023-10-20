package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest 新增团期 API请求
// taobao.alitrip.travel.fsc.route.api.project.add
//
// 新增团期
type TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest struct {
	model.Params
	// fscProjectModifyRequest
	_fscProjectModifyRequest *FscProjectModifyRequest
}

// NewTaobaoAlitripTravelFscRouteApiProjectAddRequest 初始化TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProjectAddRequest() *TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest) Reset() {
	r._fscProjectModifyRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectModifyRequest is FscProjectModifyRequest Setter
// fscProjectModifyRequest
func (r *TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest) SetFscProjectModifyRequest(_fscProjectModifyRequest *FscProjectModifyRequest) error {
	r._fscProjectModifyRequest = _fscProjectModifyRequest
	r.Set("fsc_project_modify_request", _fscProjectModifyRequest)
	return nil
}

// GetFscProjectModifyRequest FscProjectModifyRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest) GetFscProjectModifyRequest() *FscProjectModifyRequest {
	return r._fscProjectModifyRequest
}

var poolTaobaoAlitripTravelFscRouteApiProjectAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscRouteApiProjectAddRequest()
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectAddRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest
func GetTaobaoAlitripTravelFscRouteApiProjectAddAPIRequest() *TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest {
	return poolTaobaoAlitripTravelFscRouteApiProjectAddAPIRequest.Get().(*TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectAddAPIRequest 将 TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProjectAddAPIRequest(v *TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProjectAddAPIRequest.Put(v)
}
