package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDataCouponGetAPIRequest
获取优惠券信息 API请求
alibaba.data.coupon.get

获取优惠券信息，仅作客户端鉴权虚拟api使用 */
type AlibabaDataCouponGetAPIRequest struct {
	model.Params
	// 客户端鉴权虚拟api使用
	_unNamed string
}

// NewAlibabaDataCouponGetRequest 初始化AlibabaDataCouponGetAPIRequest对象
func NewAlibabaDataCouponGetRequest() *AlibabaDataCouponGetAPIRequest {
	return &AlibabaDataCouponGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDataCouponGetAPIRequest) GetApiMethodName() string {
	return "alibaba.data.coupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDataCouponGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UnNamed Setter
// 客户端鉴权虚拟api使用
func (r *AlibabaDataCouponGetAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// Get UnNamed Getter
func (r AlibabaDataCouponGetAPIRequest) GetUnNamed() string {
	return r._unNamed
}
