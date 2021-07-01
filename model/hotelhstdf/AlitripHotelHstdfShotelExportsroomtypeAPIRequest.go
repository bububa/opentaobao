package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfShotelExportsroomtypeAPIRequest
导出一个卖家房型下的所有标准房型 API请求
alitrip.hotel.hstdf.shotel.exportsroomtype

导出一个卖家酒店下的所有标准房型 */
type AlitripHotelHstdfShotelExportsroomtypeAPIRequest struct {
	model.Params
	// 卖家酒店id
	_hid int64
}

// New
