package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelMatch 酒店匹配接口-阿信
// taobao.alitrip.travel.axin.hotel.match
//
// 酒店匹配接口-阿信
func TaobaoAlitripTravelAxinHotelMatch(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelMatchAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelMatchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
