package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushvenueAPIRequest 大麦换验平台-第三方对外开放-场馆接口pushVenue API请求
// alibaba.damai.mev.open.pushvenue
//
// 开放接口推送场馆
type AlibabaDamaiMevOpenPushvenueAPIRequest struct {
	model.Params
	// 入参pushVenueParam
	_pushVenueParam *ThirdVenuePushOpenParam
}

// NewAlibabaDamaiMevOpenPushvenueRequest 初始化AlibabaDamaiMevOpenPushvenueAPIRequest对象
func NewAlibabaDamaiMevOpenPushvenueRequest() *AlibabaDamaiMevOpenPushvenueAPIRequest {
	return &AlibabaDamaiMevOpenPushvenueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushvenueAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushvenue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushvenueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenPushvenueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushVenueParam is PushVenueParam Setter
// 入参pushVenueParam
func (r *AlibabaDamaiMevOpenPushvenueAPIRequest) SetPushVenueParam(_pushVenueParam *ThirdVenuePushOpenParam) error {
	r._pushVenueParam = _pushVenueParam
	r.Set("push_venue_param", _pushVenueParam)
	return nil
}

// GetPushVenueParam PushVenueParam Getter
func (r AlibabaDamaiMevOpenPushvenueAPIRequest) GetPushVenueParam() *ThirdVenuePushOpenParam {
	return r._pushVenueParam
}
