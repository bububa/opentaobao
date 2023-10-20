package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomobilepromotionbenefitactivitysendAPIRequest 手淘专用单用户发放接口 API请求
// taobao.mobile.promotion.benefit.activity.send
//
// 卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。
type TaobaomobilepromotionbenefitactivitysendAPIRequest struct {
	model.Params
	// 单用户权益发放请求
	_singleBenefitRequest *SingleBenefitRequest
}

// NewTaobaomobilepromotionbenefitactivitysendRequest 初始化TaobaomobilepromotionbenefitactivitysendAPIRequest对象
func NewTaobaomobilepromotionbenefitactivitysendRequest() *TaobaomobilepromotionbenefitactivitysendAPIRequest {
	return &TaobaomobilepromotionbenefitactivitysendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomobilepromotionbenefitactivitysendAPIRequest) GetApiMethodName() string {
	return "taobao.mobile.promotion.benefit.activity.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomobilepromotionbenefitactivitysendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomobilepromotionbenefitactivitysendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSingleBenefitRequest is SingleBenefitRequest Setter
// 单用户权益发放请求
func (r *TaobaomobilepromotionbenefitactivitysendAPIRequest) SetSingleBenefitRequest(_singleBenefitRequest *SingleBenefitRequest) error {
	r._singleBenefitRequest = _singleBenefitRequest
	r.Set("single_benefit_request", _singleBenefitRequest)
	return nil
}

// GetSingleBenefitRequest SingleBenefitRequest Getter
func (r TaobaomobilepromotionbenefitactivitysendAPIRequest) GetSingleBenefitRequest() *SingleBenefitRequest {
	return r._singleBenefitRequest
}
