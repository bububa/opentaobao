package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelPriceBatchGet 阿信酒店批量报价查询接口
// taobao.alitrip.travel.axin.hotel.price.batch.get
//
// 阿信酒店批量报价查询接口
func TaobaoAlitripTravelAxinHotelPriceBatchGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelPriceBatchGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
