package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelRateGetmixratelistGetAPIRequest
酒店评论接口 API请求
alitrip.hotel.rate.getmixratelist.get

酒店评论接口 */
type AlitripHotelRateGetmixratelistGetAPIRequest struct {
	model.Params
	// 评论参数
	_paramGetMixRateListParam *GetMixRateListParam
}

// New
