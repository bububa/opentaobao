package hotelalliance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHmsPartnerInfoGetAPIRequest 获取合作商信息 API请求
// alitrip.hotel.hms.partner.info.get
//
// 用于给到未来酒店读取与飞猪酒店合作的合作商信息，开展单体联盟业务
type AlitripHotelHmsPartnerInfoGetAPIRequest struct {
	model.Params
	// 查询合作商信息query参数
	_queryPartnerInfoParam *QueryPartnerInfoParam
}

// NewAlitripHotelHmsPartnerInfoGetRequest 初始化AlitripHotelHmsPartnerInfoGetAPIRequest对象
func NewAlitripHotelHmsPartnerInfoGetRequest() *AlitripHotelHmsPartnerInfoGetAPIRequest {
	return &AlitripHotelHmsPartnerInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHmsPartnerInfoGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hms.partner.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHmsPartnerInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQueryPartnerInfoParam is QueryPartnerInfoParam Setter
// 查询合作商信息query参数
func (r *AlitripHotelHmsPartnerInfoGetAPIRequest) SetQueryPartnerInfoParam(_queryPartnerInfoParam *QueryPartnerInfoParam) error {
	r._queryPartnerInfoParam = _queryPartnerInfoParam
	r.Set("query_partner_info_param", _queryPartnerInfoParam)
	return nil
}

// GetQueryPartnerInfoParam QueryPartnerInfoParam Getter
func (r AlitripHotelHmsPartnerInfoGetAPIRequest) GetQueryPartnerInfoParam() *QueryPartnerInfoParam {
	return r._queryPartnerInfoParam
}
