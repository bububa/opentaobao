package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelMatch 酒店匹配接口-阿信
// taobao.alitrip.travel.axin.hotel.match
//
// 酒店匹配接口-阿信
func TaobaoAlitripTravelAxinHotelMatch(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelMatchAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelMatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
