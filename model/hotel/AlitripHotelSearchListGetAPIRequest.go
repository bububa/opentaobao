package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelSearchListGetAPIRequest
酒店搜索List接口 API请求
alitrip.hotel.search.list.get

酒店搜索List接口 */
type AlitripHotelSearchListGetAPIRequest struct {
	model.Params
	// 入参
	_paramTopHotelSearchListParam *TopHotelSearchListParam
}

// New
