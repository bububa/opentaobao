package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscHotelModifyIncrement 酒店价库变更列表查询-供销平台
// taobao.alitrip.travel.fsc.hotel.modify.increment
//
// 按照时间纬度查询酒店变更列表
func TaobaoAlitripTravelFscHotelModifyIncrement(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest, resp *axindata.TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
