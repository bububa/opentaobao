package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelCrsorderCompleteAPIRequest
CRS接送机商家服务完成接口 API请求
alitrip.travel.crsorder.complete

提供给CRS接送机商家的服务完成回调接口 */
type AlitripTravelCrsorderCompleteAPIRequest struct {
	model.Params
	// 请求对象
	_crsOrderCompleteParam *CrsOrderCompleteParam
}

// New
