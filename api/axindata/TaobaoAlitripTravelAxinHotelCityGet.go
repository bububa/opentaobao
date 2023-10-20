package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelCityGet 城市列表信息查询-阿信
// taobao.alitrip.travel.axin.hotel.city.get
//
// 阿信城市列表信息查询
func TaobaoAlitripTravelAxinHotelCityGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelCityGetAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelCityGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
