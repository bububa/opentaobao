package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelCityGet 城市列表信息查询-阿信
// taobao.alitrip.travel.axin.hotel.city.get
//
// 阿信城市列表信息查询
func TaobaoAlitripTravelAxinHotelCityGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelCityGetAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinHotelCityGetAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinHotelCityGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
