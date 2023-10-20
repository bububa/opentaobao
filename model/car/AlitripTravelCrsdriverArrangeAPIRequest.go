package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelCrsdriverArrangeAPIRequest CRS接送机商家派司机接口 API请求
// alitrip.travel.crsdriver.arrange
//
// 提供给CRS接送机商家派司机的API
type AlitripTravelCrsdriverArrangeAPIRequest struct {
	model.Params
	// 请求对象
	_crsDriverArrangeParam *CrsDriverArrangeParam
}

// NewAlitripTravelCrsdriverArrangeRequest 初始化AlitripTravelCrsdriverArrangeAPIRequest对象
func NewAlitripTravelCrsdriverArrangeRequest() *AlitripTravelCrsdriverArrangeAPIRequest {
	return &AlitripTravelCrsdriverArrangeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelCrsdriverArrangeAPIRequest) Reset() {
	r._crsDriverArrangeParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelCrsdriverArrangeAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.crsdriver.arrange"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelCrsdriverArrangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelCrsdriverArrangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrsDriverArrangeParam is CrsDriverArrangeParam Setter
// 请求对象
func (r *AlitripTravelCrsdriverArrangeAPIRequest) SetCrsDriverArrangeParam(_crsDriverArrangeParam *CrsDriverArrangeParam) error {
	r._crsDriverArrangeParam = _crsDriverArrangeParam
	r.Set("crs_driver_arrange_param", _crsDriverArrangeParam)
	return nil
}

// GetCrsDriverArrangeParam CrsDriverArrangeParam Getter
func (r AlitripTravelCrsdriverArrangeAPIRequest) GetCrsDriverArrangeParam() *CrsDriverArrangeParam {
	return r._crsDriverArrangeParam
}

var poolAlitripTravelCrsdriverArrangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelCrsdriverArrangeRequest()
	},
}

// GetAlitripTravelCrsdriverArrangeRequest 从 sync.Pool 获取 AlitripTravelCrsdriverArrangeAPIRequest
func GetAlitripTravelCrsdriverArrangeAPIRequest() *AlitripTravelCrsdriverArrangeAPIRequest {
	return poolAlitripTravelCrsdriverArrangeAPIRequest.Get().(*AlitripTravelCrsdriverArrangeAPIRequest)
}

// ReleaseAlitripTravelCrsdriverArrangeAPIRequest 将 AlitripTravelCrsdriverArrangeAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelCrsdriverArrangeAPIRequest(v *AlitripTravelCrsdriverArrangeAPIRequest) {
	v.Reset()
	poolAlitripTravelCrsdriverArrangeAPIRequest.Put(v)
}
