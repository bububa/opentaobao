package hotelalliance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelSingleInfoGetAPIRequest
获取单体酒店信息 API请求
alitrip.hotel.single.info.get

用于给到未来酒店读取与飞猪酒店合作的单体酒店信息，开展单体联盟业务 */
type AlitripHotelSingleInfoGetAPIRequest struct {
	model.Params
	// 查询酒店信息query参数
	_queryHotelInfoParam *QueryHotelInfoParam
}

// New
