package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelPriceQuery 阿信酒店分销-实时报价查询
// taobao.alitrip.travel.axin.hotel.price.query
//
// 酒店实时报价查询
func TaobaoAlitripTravelAxinHotelPriceQuery(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinHotelPriceQueryAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinHotelPriceQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
