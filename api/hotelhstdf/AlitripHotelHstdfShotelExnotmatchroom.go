package hotelhstdf

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfShotelExnotmatchroom 导出一个hid下所有未匹配rid的接口
// alitrip.hotel.hstdf.shotel.exnotmatchroom
//
// 导出一个卖家hid下所有未匹配的rid信息，包括rid，名称、英文名称、床型、窗型、面积、对外系统id
func AlitripHotelHstdfShotelExnotmatchroom(ctx context.Context, clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelExnotmatchroomAPIRequest, resp *hotelhstdf.AlitripHotelHstdfShotelExnotmatchroomAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
