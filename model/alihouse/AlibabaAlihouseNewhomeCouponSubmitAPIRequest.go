package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCouponSubmitAPIRequest 提交专车优惠券活动 API请求
// alibaba.alihouse.newhome.coupon.submit
//
// 提交专车优惠券活动
type AlibabaAlihouseNewhomeCouponSubmitAPIRequest struct {
	model.Params
	// 打车券活动对象
	_couponDto *MarketingCouponDto
}

// NewAlibabaAlihouseNewhomeCouponSubmitRequest 初始化AlibabaAlihouseNewhomeCouponSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeCouponSubmitRequest() *AlibabaAlihouseNewhomeCouponSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeCouponSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeCouponSubmitAPIRequest) Reset() {
	r._couponDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.coupon.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCouponDto is CouponDto Setter
// 打车券活动对象
func (r *AlibabaAlihouseNewhomeCouponSubmitAPIRequest) SetCouponDto(_couponDto *MarketingCouponDto) error {
	r._couponDto = _couponDto
	r.Set("coupon_dto", _couponDto)
	return nil
}

// GetCouponDto CouponDto Getter
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetCouponDto() *MarketingCouponDto {
	return r._couponDto
}

var poolAlibabaAlihouseNewhomeCouponSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeCouponSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeCouponSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeCouponSubmitAPIRequest
func GetAlibabaAlihouseNewhomeCouponSubmitAPIRequest() *AlibabaAlihouseNewhomeCouponSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeCouponSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeCouponSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeCouponSubmitAPIRequest 将 AlibabaAlihouseNewhomeCouponSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeCouponSubmitAPIRequest(v *AlibabaAlihouseNewhomeCouponSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeCouponSubmitAPIRequest.Put(v)
}
