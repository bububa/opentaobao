package hotelalliance

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelalliance"
)

// AlitripHotelHmsPartnerInfoGet 获取合作商信息
// alitrip.hotel.hms.partner.info.get
//
// 用于给到未来酒店读取与飞猪酒店合作的合作商信息，开展单体联盟业务
func AlitripHotelHmsPartnerInfoGet(ctx context.Context, clt *core.SDKClient, req *hotelalliance.AlitripHotelHmsPartnerInfoGetAPIRequest, resp *hotelalliance.AlitripHotelHmsPartnerInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
