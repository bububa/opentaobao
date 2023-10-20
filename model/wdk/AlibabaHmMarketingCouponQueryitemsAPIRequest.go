package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingCouponQueryitemsAPIRequest 查询优惠券活动下的商品 API请求
// alibaba.hm.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
type AlibabaHmMarketingCouponQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaHmMarketingCouponQueryitemsRequest 初始化AlibabaHmMarketingCouponQueryitemsAPIRequest对象
func NewAlibabaHmMarketingCouponQueryitemsRequest() *AlibabaHmMarketingCouponQueryitemsAPIRequest {
	return &AlibabaHmMarketingCouponQueryitemsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingCouponQueryitemsAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingCouponQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.coupon.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingCouponQueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingCouponQueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaHmMarketingCouponQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingCouponQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}

var poolAlibabaHmMarketingCouponQueryitemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingCouponQueryitemsRequest()
	},
}

// GetAlibabaHmMarketingCouponQueryitemsRequest 从 sync.Pool 获取 AlibabaHmMarketingCouponQueryitemsAPIRequest
func GetAlibabaHmMarketingCouponQueryitemsAPIRequest() *AlibabaHmMarketingCouponQueryitemsAPIRequest {
	return poolAlibabaHmMarketingCouponQueryitemsAPIRequest.Get().(*AlibabaHmMarketingCouponQueryitemsAPIRequest)
}

// ReleaseAlibabaHmMarketingCouponQueryitemsAPIRequest 将 AlibabaHmMarketingCouponQueryitemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingCouponQueryitemsAPIRequest(v *AlibabaHmMarketingCouponQueryitemsAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingCouponQueryitemsAPIRequest.Put(v)
}
