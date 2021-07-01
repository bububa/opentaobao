package hotelalliance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHmsPartnerInfoGetAPIRequest
获取合作商信息 API请求
alitrip.hotel.hms.partner.info.get

用于给到未来酒店读取与飞猪酒店合作的合作商信息，开展单体联盟业务 */
type AlitripHotelHmsPartnerInfoGetAPIRequest struct {
	model.Params
	// 查询合作商信息query参数
	_queryPartnerInfoParam *QueryPartnerInfoParam
}

// New
