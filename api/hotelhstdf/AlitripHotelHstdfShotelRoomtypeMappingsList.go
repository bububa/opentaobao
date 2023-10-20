package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfShotelRoomtypeMappingsList 根据HID获取所有卖家房型匹配关系
// alitrip.hotel.hstdf.shotel.roomtype.mappings.list
//
// 根据HID获取所有卖家房型匹配关系
func AlitripHotelHstdfShotelRoomtypeMappingsList(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest, resp *hotelhstdf.AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
