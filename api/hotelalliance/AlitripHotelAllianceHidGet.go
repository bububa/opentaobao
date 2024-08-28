package hotelalliance

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelalliance"
)

// AlitripHotelAllianceHidGet 获取联盟hid
// alitrip.hotel.alliance.hid.get
//
// 获取符合条件的菲住联盟hid，目前支持指定日期上线的菲住联盟hid查询
func AlitripHotelAllianceHidGet(ctx context.Context, clt *core.SDKClient, req *hotelalliance.AlitripHotelAllianceHidGetAPIRequest, resp *hotelalliance.AlitripHotelAllianceHidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
