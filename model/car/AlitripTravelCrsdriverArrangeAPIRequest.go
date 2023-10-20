package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelcrsdriverarrangeAPIRequest CRS接送机商家派司机接口 API请求
// alitrip.travel.crsdriver.arrange
//
// 提供给CRS接送机商家派司机的API
type AlitriptravelcrsdriverarrangeAPIRequest struct {
	model.Params
	// 请求对象
	_crsDriverArrangeParam *CrsDriverArrangeParam
}

// NewAlitriptravelcrsdriverarrangeRequest 初始化AlitriptravelcrsdriverarrangeAPIRequest对象
func NewAlitriptravelcrsdriverarrangeRequest() *AlitriptravelcrsdriverarrangeAPIRequest {
	return &AlitriptravelcrsdriverarrangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelcrsdriverarrangeAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.crsdriver.arrange"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelcrsdriverarrangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelcrsdriverarrangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrsDriverArrangeParam is CrsDriverArrangeParam Setter
// 请求对象
func (r *AlitriptravelcrsdriverarrangeAPIRequest) SetCrsDriverArrangeParam(_crsDriverArrangeParam *CrsDriverArrangeParam) error {
	r._crsDriverArrangeParam = _crsDriverArrangeParam
	r.Set("crs_driver_arrange_param", _crsDriverArrangeParam)
	return nil
}

// GetCrsDriverArrangeParam CrsDriverArrangeParam Getter
func (r AlitriptravelcrsdriverarrangeAPIRequest) GetCrsDriverArrangeParam() *CrsDriverArrangeParam {
	return r._crsDriverArrangeParam
}
