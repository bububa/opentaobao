package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelCrsorderCompleteAPIRequest CRS接送机商家服务完成接口 API请求
// alitrip.travel.crsorder.complete
//
// 提供给CRS接送机商家的服务完成回调接口
type AlitripTravelCrsorderCompleteAPIRequest struct {
	model.Params
	// 请求对象
	_crsOrderCompleteParam *CrsOrderCompleteParam
}

// NewAlitripTravelCrsorderCompleteRequest 初始化AlitripTravelCrsorderCompleteAPIRequest对象
func NewAlitripTravelCrsorderCompleteRequest() *AlitripTravelCrsorderCompleteAPIRequest {
	return &AlitripTravelCrsorderCompleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelCrsorderCompleteAPIRequest) Reset() {
	r._crsOrderCompleteParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelCrsorderCompleteAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.crsorder.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelCrsorderCompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelCrsorderCompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrsOrderCompleteParam is CrsOrderCompleteParam Setter
// 请求对象
func (r *AlitripTravelCrsorderCompleteAPIRequest) SetCrsOrderCompleteParam(_crsOrderCompleteParam *CrsOrderCompleteParam) error {
	r._crsOrderCompleteParam = _crsOrderCompleteParam
	r.Set("crs_order_complete_param", _crsOrderCompleteParam)
	return nil
}

// GetCrsOrderCompleteParam CrsOrderCompleteParam Getter
func (r AlitripTravelCrsorderCompleteAPIRequest) GetCrsOrderCompleteParam() *CrsOrderCompleteParam {
	return r._crsOrderCompleteParam
}

var poolAlitripTravelCrsorderCompleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelCrsorderCompleteRequest()
	},
}

// GetAlitripTravelCrsorderCompleteRequest 从 sync.Pool 获取 AlitripTravelCrsorderCompleteAPIRequest
func GetAlitripTravelCrsorderCompleteAPIRequest() *AlitripTravelCrsorderCompleteAPIRequest {
	return poolAlitripTravelCrsorderCompleteAPIRequest.Get().(*AlitripTravelCrsorderCompleteAPIRequest)
}

// ReleaseAlitripTravelCrsorderCompleteAPIRequest 将 AlitripTravelCrsorderCompleteAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelCrsorderCompleteAPIRequest(v *AlitripTravelCrsorderCompleteAPIRequest) {
	v.Reset()
	poolAlitripTravelCrsorderCompleteAPIRequest.Put(v)
}
