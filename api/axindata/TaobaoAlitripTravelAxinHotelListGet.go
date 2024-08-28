package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelListGet 标准酒店信息查询-阿信
// taobao.alitrip.travel.axin.hotel.list.get
//
// 阿信酒店分销基础数据查询
func TaobaoAlitripTravelAxinHotelListGet(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelListGetAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
