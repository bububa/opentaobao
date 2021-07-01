package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelDetailStaticinfoGetAPIRequest
详情页静态信息 API请求
alitrip.hotel.detail.staticinfo.get

酒店静态信息TOP接口 */
type AlitripHotelDetailStaticinfoGetAPIRequest struct {
	model.Params
	// 酒店详情页静态信息参数
	_paramHotelTopStaticDetailsParam *HotelTopStaticDetailsParam
}

// New
