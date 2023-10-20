package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfShotelMatchshotelself 自主匹配标准酒店以及卖家酒店
// alitrip.hotel.hstdf.shotel.matchshotelself
//
// 商家通过指定的标准酒店id和卖家酒店id进行匹配
func AlitripHotelHstdfShotelMatchshotelself(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelMatchshotelselfAPIRequest, resp *hotelhstdf.AlitripHotelHstdfShotelMatchshotelselfAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
