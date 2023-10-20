package eleenterprisecoupon

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseCouponGetAPIRequest 获取用户优惠券 API请求
// alibaba.ele.enterprise.coupon.get
//
// 获取用户优惠券
type AlibabaEleEnterpriseCouponGetAPIRequest struct {
	model.Params
	// 手机号
	_phone string
}

// NewAlibabaEleEnterpriseCouponGetRequest 初始化AlibabaEleEnterpriseCouponGetAPIRequest对象
func NewAlibabaEleEnterpriseCouponGetRequest() *AlibabaEleEnterpriseCouponGetAPIRequest {
	return &AlibabaEleEnterpriseCouponGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseCouponGetAPIRequest) Reset() {
	r._phone = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCouponGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.coupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCouponGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseCouponGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *AlibabaEleEnterpriseCouponGetAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaEleEnterpriseCouponGetAPIRequest) GetPhone() string {
	return r._phone
}

var poolAlibabaEleEnterpriseCouponGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseCouponGetRequest()
	},
}

// GetAlibabaEleEnterpriseCouponGetRequest 从 sync.Pool 获取 AlibabaEleEnterpriseCouponGetAPIRequest
func GetAlibabaEleEnterpriseCouponGetAPIRequest() *AlibabaEleEnterpriseCouponGetAPIRequest {
	return poolAlibabaEleEnterpriseCouponGetAPIRequest.Get().(*AlibabaEleEnterpriseCouponGetAPIRequest)
}

// ReleaseAlibabaEleEnterpriseCouponGetAPIRequest 将 AlibabaEleEnterpriseCouponGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseCouponGetAPIRequest(v *AlibabaEleEnterpriseCouponGetAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseCouponGetAPIRequest.Put(v)
}
