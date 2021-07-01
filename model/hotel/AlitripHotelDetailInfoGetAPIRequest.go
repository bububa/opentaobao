package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelDetailInfoGetAPIRequest
详情页动态信息接口 API请求
alitrip.hotel.detail.info.get

酒店详情页动态信息TOP方法 */
type AlitripHotelDetailInfoGetAPIRequest struct {
	model.Params
	// 详情页动态信息参数类
	_paramHotelTopDetailsParam *HotelTopDetailsParam
}

// New
