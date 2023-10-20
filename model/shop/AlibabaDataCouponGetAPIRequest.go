package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDataCouponGetAPIRequest 获取优惠券信息 API请求
// alibaba.data.coupon.get
//
// 获取优惠券信息，仅作客户端鉴权虚拟api使用
type AlibabaDataCouponGetAPIRequest struct {
	model.Params
	// 客户端鉴权虚拟api使用
	_unNamed string
}

// NewAlibabaDataCouponGetRequest 初始化AlibabaDataCouponGetAPIRequest对象
func NewAlibabaDataCouponGetRequest() *AlibabaDataCouponGetAPIRequest {
	return &AlibabaDataCouponGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDataCouponGetAPIRequest) Reset() {
	r._unNamed = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDataCouponGetAPIRequest) GetApiMethodName() string {
	return "alibaba.data.coupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDataCouponGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDataCouponGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnNamed is UnNamed Setter
// 客户端鉴权虚拟api使用
func (r *AlibabaDataCouponGetAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// GetUnNamed UnNamed Getter
func (r AlibabaDataCouponGetAPIRequest) GetUnNamed() string {
	return r._unNamed
}

var poolAlibabaDataCouponGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDataCouponGetRequest()
	},
}

// GetAlibabaDataCouponGetRequest 从 sync.Pool 获取 AlibabaDataCouponGetAPIRequest
func GetAlibabaDataCouponGetAPIRequest() *AlibabaDataCouponGetAPIRequest {
	return poolAlibabaDataCouponGetAPIRequest.Get().(*AlibabaDataCouponGetAPIRequest)
}

// ReleaseAlibabaDataCouponGetAPIRequest 将 AlibabaDataCouponGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaDataCouponGetAPIRequest(v *AlibabaDataCouponGetAPIRequest) {
	v.Reset()
	poolAlibabaDataCouponGetAPIRequest.Put(v)
}
