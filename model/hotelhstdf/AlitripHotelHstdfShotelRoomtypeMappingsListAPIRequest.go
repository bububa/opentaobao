package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest
根据HID获取所有卖家房型匹配关系 API请求
alitrip.hotel.hstdf.shotel.roomtype.mappings.list

根据HID获取所有卖家房型匹配关系 */
type AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest struct {
	model.Params
	// HID
	_hid int64
}

// New
