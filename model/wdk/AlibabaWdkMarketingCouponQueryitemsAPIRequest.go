package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingcouponqueryitemsAPIRequest 查询优惠券活动下的商品 API请求
// alibaba.wdk.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
type AlibabawdkmarketingcouponqueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabawdkmarketingcouponqueryitemsRequest 初始化AlibabawdkmarketingcouponqueryitemsAPIRequest对象
func NewAlibabawdkmarketingcouponqueryitemsRequest() *AlibabawdkmarketingcouponqueryitemsAPIRequest {
	return &AlibabawdkmarketingcouponqueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingcouponqueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingcouponqueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingcouponqueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabawdkmarketingcouponqueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingcouponqueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
