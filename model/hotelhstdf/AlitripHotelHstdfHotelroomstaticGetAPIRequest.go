package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfHotelroomstaticGetAPIRequest
根据类型查询静态字段 API请求
alitrip.hotel.hstdf.hotelroomstatic.get

根据类型查询分页静态字段 */
type AlitripHotelHstdfHotelroomstaticGetAPIRequest struct {
	model.Params
	// 参数封装
	_paramGetHotelRoomStaticParam *GetHotelRoomStaticParam
}

// New
