package hotelhstdf

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfShotelExportsroomtype 导出一个卖家房型下的所有标准房型
// alitrip.hotel.hstdf.shotel.exportsroomtype
//
// 导出一个卖家酒店下的所有标准房型
func AlitripHotelHstdfShotelExportsroomtype(ctx context.Context, clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelExportsroomtypeAPIRequest, resp *hotelhstdf.AlitripHotelHstdfShotelExportsroomtypeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
