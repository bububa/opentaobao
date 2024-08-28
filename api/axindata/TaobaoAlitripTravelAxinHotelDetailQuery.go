package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelDetailQuery 阿信酒店分销-标准酒店详情查询
// taobao.alitrip.travel.axin.hotel.detail.query
//
// 标准酒店详情查询
func TaobaoAlitripTravelAxinHotelDetailQuery(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelDetailQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
