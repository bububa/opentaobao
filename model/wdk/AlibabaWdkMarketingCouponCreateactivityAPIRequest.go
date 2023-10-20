package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingcouponcreateactivityAPIRequest 优惠券活动创建 API请求
// alibaba.wdk.marketing.coupon.createactivity
//
// 添加优惠券活动
type AlibabawdkmarketingcouponcreateactivityAPIRequest struct {
	model.Params
	// 创建优惠券活动请求入参
	_param *CouponActivity
}

// NewAlibabawdkmarketingcouponcreateactivityRequest 初始化AlibabawdkmarketingcouponcreateactivityAPIRequest对象
func NewAlibabawdkmarketingcouponcreateactivityRequest() *AlibabawdkmarketingcouponcreateactivityAPIRequest {
	return &AlibabawdkmarketingcouponcreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingcouponcreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingcouponcreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingcouponcreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建优惠券活动请求入参
func (r *AlibabawdkmarketingcouponcreateactivityAPIRequest) SetParam(_param *CouponActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingcouponcreateactivityAPIRequest) GetParam() *CouponActivity {
	return r._param
}
