package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingcouponqueryitemsAPIRequest 查询优惠券活动下的商品 API请求
// alibaba.hm.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
type AlibabahmmarketingcouponqueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabahmmarketingcouponqueryitemsRequest 初始化AlibabahmmarketingcouponqueryitemsAPIRequest对象
func NewAlibabahmmarketingcouponqueryitemsRequest() *AlibabahmmarketingcouponqueryitemsAPIRequest {
	return &AlibabahmmarketingcouponqueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingcouponqueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.coupon.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingcouponqueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingcouponqueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabahmmarketingcouponqueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingcouponqueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
