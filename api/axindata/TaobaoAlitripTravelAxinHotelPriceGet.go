package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelPriceGet 酒店报价服务-阿信
// taobao.alitrip.travel.axin.hotel.price.get
//
// 酒店报价查询服务
func TaobaoAlitripTravelAxinHotelPriceGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelPriceGetAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelPriceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
