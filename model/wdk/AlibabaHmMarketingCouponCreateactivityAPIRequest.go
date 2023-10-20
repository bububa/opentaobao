package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingcouponcreateactivityAPIRequest 优惠券活动创建 API请求
// alibaba.hm.marketing.coupon.createactivity
//
// 添加优惠券活动
type AlibabahmmarketingcouponcreateactivityAPIRequest struct {
	model.Params
	// 创建优惠券活动请求入参
	_param *CouponActivity
}

// NewAlibabahmmarketingcouponcreateactivityRequest 初始化AlibabahmmarketingcouponcreateactivityAPIRequest对象
func NewAlibabahmmarketingcouponcreateactivityRequest() *AlibabahmmarketingcouponcreateactivityAPIRequest {
	return &AlibabahmmarketingcouponcreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingcouponcreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.coupon.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingcouponcreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingcouponcreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建优惠券活动请求入参
func (r *AlibabahmmarketingcouponcreateactivityAPIRequest) SetParam(_param *CouponActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingcouponcreateactivityAPIRequest) GetParam() *CouponActivity {
	return r._param
}
