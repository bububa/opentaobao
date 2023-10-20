package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponQueryitemsAPIRequest 查询优惠券活动下的商品 API请求
// alibaba.wdk.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
type AlibabaWdkMarketingCouponQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaWdkMarketingCouponQueryitemsRequest 初始化AlibabaWdkMarketingCouponQueryitemsAPIRequest对象
func NewAlibabaWdkMarketingCouponQueryitemsRequest() *AlibabaWdkMarketingCouponQueryitemsAPIRequest {
	return &AlibabaWdkMarketingCouponQueryitemsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingCouponQueryitemsAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponQueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingCouponQueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaWdkMarketingCouponQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingCouponQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}

var poolAlibabaWdkMarketingCouponQueryitemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingCouponQueryitemsRequest()
	},
}

// GetAlibabaWdkMarketingCouponQueryitemsRequest 从 sync.Pool 获取 AlibabaWdkMarketingCouponQueryitemsAPIRequest
func GetAlibabaWdkMarketingCouponQueryitemsAPIRequest() *AlibabaWdkMarketingCouponQueryitemsAPIRequest {
	return poolAlibabaWdkMarketingCouponQueryitemsAPIRequest.Get().(*AlibabaWdkMarketingCouponQueryitemsAPIRequest)
}

// ReleaseAlibabaWdkMarketingCouponQueryitemsAPIRequest 将 AlibabaWdkMarketingCouponQueryitemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingCouponQueryitemsAPIRequest(v *AlibabaWdkMarketingCouponQueryitemsAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingCouponQueryitemsAPIRequest.Put(v)
}
