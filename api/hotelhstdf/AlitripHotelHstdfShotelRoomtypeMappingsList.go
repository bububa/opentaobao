package hotelhstdf

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfShotelRoomtypeMappingsList 根据HID获取所有卖家房型匹配关系
// alitrip.hotel.hstdf.shotel.roomtype.mappings.list
//
// 根据HID获取所有卖家房型匹配关系
func AlitripHotelHstdfShotelRoomtypeMappingsList(ctx context.Context, clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest, resp *hotelhstdf.AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
