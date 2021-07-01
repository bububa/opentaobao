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

// NewAlitripHotelSingleInfoGetRequest 初始化AlitripHotelSingleInfoGetAPIRequest对象
func NewAlitripHotelSingleInfoGetRequest() *AlitripHotelSingleInfoGetAPIRequest {
	return &AlitripHotelSingleInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelSingleInfoGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.single.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelSingleInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is QueryHotelInfoParam Setter
// 查询酒店信息query参数
func (r *AlitripHotelSingleInfoGetAPIRequest) SetQueryHotelInfoParam(_queryHotelInfoParam *QueryHotelInfoParam) error {
	r._queryHotelInfoParam = _queryHotelInfoParam
	r.Set("query_hotel_info_param", _queryHotelInfoParam)
	return nil
}

// Get QueryHotelInfoParam Getter
func (r AlitripHotelSingleInfoGetAPIRequest) GetQueryHotelInfoParam() *QueryHotelInfoParam {
	return r._queryHotelInfoParam
}
