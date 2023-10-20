package hotelalliance

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripHotelHmsPartnerInfoGetAPIRequest) Reset() {
	r._queryPartnerInfoParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHmsPartnerInfoGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hms.partner.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHmsPartnerInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelHmsPartnerInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripHotelHmsPartnerInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripHotelHmsPartnerInfoGetRequest()
	},
}

// GetAlitripHotelHmsPartnerInfoGetRequest 从 sync.Pool 获取 AlitripHotelHmsPartnerInfoGetAPIRequest
func GetAlitripHotelHmsPartnerInfoGetAPIRequest() *AlitripHotelHmsPartnerInfoGetAPIRequest {
	return poolAlitripHotelHmsPartnerInfoGetAPIRequest.Get().(*AlitripHotelHmsPartnerInfoGetAPIRequest)
}

// ReleaseAlitripHotelHmsPartnerInfoGetAPIRequest 将 AlitripHotelHmsPartnerInfoGetAPIRequest 放入 sync.Pool
func ReleaseAlitripHotelHmsPartnerInfoGetAPIRequest(v *AlitripHotelHmsPartnerInfoGetAPIRequest) {
	v.Reset()
	poolAlitripHotelHmsPartnerInfoGetAPIRequest.Put(v)
}
