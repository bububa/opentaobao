package hotelalliance

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelalliance"
)

// AlitripHotelSingleInfoGet 获取单体酒店信息
// alitrip.hotel.single.info.get
//
// 用于给到未来酒店读取与飞猪酒店合作的单体酒店信息，开展单体联盟业务
func AlitripHotelSingleInfoGet(ctx context.Context, clt *core.SDKClient, req *hotelalliance.AlitripHotelSingleInfoGetAPIRequest, resp *hotelalliance.AlitripHotelSingleInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
