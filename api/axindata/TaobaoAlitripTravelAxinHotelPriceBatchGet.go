package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelPriceBatchGet 阿信酒店批量报价查询接口
// taobao.alitrip.travel.axin.hotel.price.batch.get
//
// 阿信酒店批量报价查询接口
func TaobaoAlitripTravelAxinHotelPriceBatchGet(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelPriceBatchGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
