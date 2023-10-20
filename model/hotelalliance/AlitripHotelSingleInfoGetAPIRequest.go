package hotelalliance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelsingleinfogetAPIRequest 获取单体酒店信息 API请求
// alitrip.hotel.single.info.get
//
// 用于给到未来酒店读取与飞猪酒店合作的单体酒店信息，开展单体联盟业务
type AlitriphotelsingleinfogetAPIRequest struct {
	model.Params
	// 查询酒店信息query参数
	_queryHotelInfoParam *QueryHotelInfoParam
}

// NewAlitriphotelsingleinfogetRequest 初始化AlitriphotelsingleinfogetAPIRequest对象
func NewAlitriphotelsingleinfogetRequest() *AlitriphotelsingleinfogetAPIRequest {
	return &AlitriphotelsingleinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelsingleinfogetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.single.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelsingleinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelsingleinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryHotelInfoParam is QueryHotelInfoParam Setter
// 查询酒店信息query参数
func (r *AlitriphotelsingleinfogetAPIRequest) SetQueryHotelInfoParam(_queryHotelInfoParam *QueryHotelInfoParam) error {
	r._queryHotelInfoParam = _queryHotelInfoParam
	r.Set("query_hotel_info_param", _queryHotelInfoParam)
	return nil
}

// GetQueryHotelInfoParam QueryHotelInfoParam Getter
func (r AlitriphotelsingleinfogetAPIRequest) GetQueryHotelInfoParam() *QueryHotelInfoParam {
	return r._queryHotelInfoParam
}
