package hotelhstdf

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfShotelMatchsroomself 匹配标准房型以及卖家房型
// alitrip.hotel.hstdf.shotel.matchsroomself
//
// 匹配卖家房型以及标准房型，返回匹配结果
func AlitripHotelHstdfShotelMatchsroomself(ctx context.Context, clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelMatchsroomselfAPIRequest, resp *hotelhstdf.AlitripHotelHstdfShotelMatchsroomselfAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
