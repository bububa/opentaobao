package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelSearchListGetAPIRequest
酒店搜索List接口 API请求
alitrip.hotel.search.list.get

酒店搜索List接口 */
type AlitripHotelSearchListGetAPIRequest struct {
	model.Params
	// 入参
	_paramTopHotelSearchListParam *TopHotelSearchListParam
}

// NewAlitripHotelSearchListGetRequest 初始化AlitripHotelSearchListGetAPIRequest对象
func NewAlitripHotelSearchListGetRequest() *AlitripHotelSearchListGetAPIRequest {
	return &AlitripHotelSearchListGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelSearchListGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.search.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelSearchListGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamTopHotelSearchListParam Setter
// 入参
func (r *AlitripHotelSearchListGetAPIRequest) SetParamTopHotelSearchListParam(_paramTopHotelSearchListParam *TopHotelSearchListParam) error {
	r._paramTopHotelSearchListParam = _paramTopHotelSearchListParam
	r.Set("param_top_hotel_search_list_param", _paramTopHotelSearchListParam)
	return nil
}

// Get ParamTopHotelSearchListParam Getter
func (r AlitripHotelSearchListGetAPIRequest) GetParamTopHotelSearchListParam() *TopHotelSearchListParam {
	return r._paramTopHotelSearchListParam
}
