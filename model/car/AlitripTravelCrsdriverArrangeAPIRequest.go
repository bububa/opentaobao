package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelCrsdriverArrangeAPIRequest
CRS接送机商家派司机接口 API请求
alitrip.travel.crsdriver.arrange

提供给CRS接送机商家派司机的API */
type AlitripTravelCrsdriverArrangeAPIRequest struct {
	model.Params
	// 请求对象
	_crsDriverArrangeParam *CrsDriverArrangeParam
}

// NewAlitripTravelCrsdriverArrangeRequest 初始化AlitripTravelCrsdriverArrangeAPIRequest对象
func NewAlitripTravelCrsdriverArrangeRequest() *AlitripTravelCrsdriverArrangeAPIRequest {
	return &AlitripTravelCrsdriverArrangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelCrsdriverArrangeAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.crsdriver.arrange"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelCrsdriverArrangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CrsDriverArrangeParam Setter
// 请求对象
func (r *AlitripTravelCrsdriverArrangeAPIRequest) SetCrsDriverArrangeParam(_crsDriverArrangeParam *CrsDriverArrangeParam) error {
	r._crsDriverArrangeParam = _crsDriverArrangeParam
	r.Set("crs_driver_arrange_param", _crsDriverArrangeParam)
	return nil
}

// Get CrsDriverArrangeParam Getter
func (r AlitripTravelCrsdriverArrangeAPIRequest) GetCrsDriverArrangeParam() *CrsDriverArrangeParam {
	return r._crsDriverArrangeParam
}
