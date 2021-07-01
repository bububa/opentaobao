package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfShotelMatchshotelselfAPIRequest
自主匹配标准酒店以及卖家酒店 API请求
alitrip.hotel.hstdf.shotel.matchshotelself

商家通过指定的标准酒店id和卖家酒店id进行匹配 */
type AlitripHotelHstdfShotelMatchshotelselfAPIRequest struct {
	model.Params
	// HotelMatchParam
	_param0 *HotelMatchParam
}

// New
