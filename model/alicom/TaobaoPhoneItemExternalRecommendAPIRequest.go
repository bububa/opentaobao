package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaophoneitemexternalrecommendAPIRequest 话费选品能力外放 API请求
// taobao.phone.item.external.recommend
//
// 话费选品能力外放
type TaobaophoneitemexternalrecommendAPIRequest struct {
	model.Params
	// 选品请求
	_phoneDistributionRecommendReq *PhoneDistributionRecommendReq
}

// NewTaobaophoneitemexternalrecommendRequest 初始化TaobaophoneitemexternalrecommendAPIRequest对象
func NewTaobaophoneitemexternalrecommendRequest() *TaobaophoneitemexternalrecommendAPIRequest {
	return &TaobaophoneitemexternalrecommendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaophoneitemexternalrecommendAPIRequest) GetApiMethodName() string {
	return "taobao.phone.item.external.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaophoneitemexternalrecommendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaophoneitemexternalrecommendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhoneDistributionRecommendReq is PhoneDistributionRecommendReq Setter
// 选品请求
func (r *TaobaophoneitemexternalrecommendAPIRequest) SetPhoneDistributionRecommendReq(_phoneDistributionRecommendReq *PhoneDistributionRecommendReq) error {
	r._phoneDistributionRecommendReq = _phoneDistributionRecommendReq
	r.Set("phone_distribution_recommend_req", _phoneDistributionRecommendReq)
	return nil
}

// GetPhoneDistributionRecommendReq PhoneDistributionRecommendReq Getter
func (r TaobaophoneitemexternalrecommendAPIRequest) GetPhoneDistributionRecommendReq() *PhoneDistributionRecommendReq {
	return r._phoneDistributionRecommendReq
}
