package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfShotelMatchsroomselfAPIRequest
匹配标准房型以及卖家房型 API请求
alitrip.hotel.hstdf.shotel.matchsroomself

匹配卖家房型以及标准房型，返回匹配结果 */
type AlitripHotelHstdfShotelMatchsroomselfAPIRequest struct {
	model.Params
	// SroomTypeMatchParam
	_param0 *SroomTypeMatchParam
}

// New
