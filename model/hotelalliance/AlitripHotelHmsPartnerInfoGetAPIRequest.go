package hotelalliance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhmspartnerinfogetAPIRequest 获取合作商信息 API请求
// alitrip.hotel.hms.partner.info.get
//
// 用于给到未来酒店读取与飞猪酒店合作的合作商信息，开展单体联盟业务
type AlitriphotelhmspartnerinfogetAPIRequest struct {
	model.Params
	// 查询合作商信息query参数
	_queryPartnerInfoParam *QueryPartnerInfoParam
}

// NewAlitriphotelhmspartnerinfogetRequest 初始化AlitriphotelhmspartnerinfogetAPIRequest对象
func NewAlitriphotelhmspartnerinfogetRequest() *AlitriphotelhmspartnerinfogetAPIRequest {
	return &AlitriphotelhmspartnerinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelhmspartnerinfogetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hms.partner.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelhmspartnerinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelhmspartnerinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryPartnerInfoParam is QueryPartnerInfoParam Setter
// 查询合作商信息query参数
func (r *AlitriphotelhmspartnerinfogetAPIRequest) SetQueryPartnerInfoParam(_queryPartnerInfoParam *QueryPartnerInfoParam) error {
	r._queryPartnerInfoParam = _queryPartnerInfoParam
	r.Set("query_partner_info_param", _queryPartnerInfoParam)
	return nil
}

// GetQueryPartnerInfoParam QueryPartnerInfoParam Getter
func (r AlitriphotelhmspartnerinfogetAPIRequest) GetQueryPartnerInfoParam() *QueryPartnerInfoParam {
	return r._queryPartnerInfoParam
}
