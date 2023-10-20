package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractCouponApplyAPIRequest 优惠券领取鉴权接口 API请求
// alibaba.interact.coupon.apply
//
// 鉴权接口，为coupon.apply接口鉴权
type AlibabaInteractCouponApplyAPIRequest struct {
	model.Params
}

// NewAlibabaInteractCouponApplyRequest 初始化AlibabaInteractCouponApplyAPIRequest对象
func NewAlibabaInteractCouponApplyRequest() *AlibabaInteractCouponApplyAPIRequest {
	return &AlibabaInteractCouponApplyAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractCouponApplyAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractCouponApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractCouponApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractCouponApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractCouponApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractCouponApplyRequest()
	},
}

// GetAlibabaInteractCouponApplyRequest 从 sync.Pool 获取 AlibabaInteractCouponApplyAPIRequest
func GetAlibabaInteractCouponApplyAPIRequest() *AlibabaInteractCouponApplyAPIRequest {
	return poolAlibabaInteractCouponApplyAPIRequest.Get().(*AlibabaInteractCouponApplyAPIRequest)
}

// ReleaseAlibabaInteractCouponApplyAPIRequest 将 AlibabaInteractCouponApplyAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractCouponApplyAPIRequest(v *AlibabaInteractCouponApplyAPIRequest) {
	v.Reset()
	poolAlibabaInteractCouponApplyAPIRequest.Put(v)
}
