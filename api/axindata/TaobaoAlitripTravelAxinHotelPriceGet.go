package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelPriceGet 酒店报价服务-阿信
// taobao.alitrip.travel.axin.hotel.price.get
//
// 酒店报价查询服务
func TaobaoAlitripTravelAxinHotelPriceGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelPriceGetAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinHotelPriceGetAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinHotelPriceGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
