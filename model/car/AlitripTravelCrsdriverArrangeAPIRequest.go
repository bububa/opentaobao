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

// New
