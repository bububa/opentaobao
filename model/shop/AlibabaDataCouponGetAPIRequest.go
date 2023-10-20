package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadatacoupongetAPIRequest 获取优惠券信息 API请求
// alibaba.data.coupon.get
//
// 获取优惠券信息，仅作客户端鉴权虚拟api使用
type AlibabadatacoupongetAPIRequest struct {
	model.Params
	// 客户端鉴权虚拟api使用
	_unNamed string
}

// NewAlibabadatacoupongetRequest 初始化AlibabadatacoupongetAPIRequest对象
func NewAlibabadatacoupongetRequest() *AlibabadatacoupongetAPIRequest {
	return &AlibabadatacoupongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadatacoupongetAPIRequest) GetApiMethodName() string {
	return "alibaba.data.coupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadatacoupongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadatacoupongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnNamed is UnNamed Setter
// 客户端鉴权虚拟api使用
func (r *AlibabadatacoupongetAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// GetUnNamed UnNamed Getter
func (r AlibabadatacoupongetAPIRequest) GetUnNamed() string {
	return r._unNamed
}
