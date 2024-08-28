package hotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// AlitripHotelRateGetmixratelistGet 酒店评论接口
// alitrip.hotel.rate.getmixratelist.get
//
// 酒店评论接口
func AlitripHotelRateGetmixratelistGet(ctx context.Context, clt *core.SDKClient, req *hotel.AlitripHotelRateGetmixratelistGetAPIRequest, resp *hotel.AlitripHotelRateGetmixratelistGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
