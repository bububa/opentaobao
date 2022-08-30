package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelPriceBatchGet 阿信酒店批量报价查询接口
// taobao.alitrip.travel.axin.hotel.price.batch.get
//
// 阿信酒店批量报价查询接口
func TaobaoAlitripTravelAxinHotelPriceBatchGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinHotelPriceBatchGetAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinHotelPriceBatchGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
