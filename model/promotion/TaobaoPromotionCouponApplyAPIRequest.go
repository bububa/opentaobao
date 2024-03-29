package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponApplyAPIRequest 优惠券领取 API请求
// taobao.promotion.coupon.apply
//
// 优惠券领取
type TaobaoPromotionCouponApplyAPIRequest struct {
	model.Params
	// 卖家id
	_sellerId string
	// 传播id
	_spreadId string
}

// NewTaobaoPromotionCouponApplyRequest 初始化TaobaoPromotionCouponApplyAPIRequest对象
func NewTaobaoPromotionCouponApplyRequest() *TaobaoPromotionCouponApplyAPIRequest {
	return &TaobaoPromotionCouponApplyAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionCouponApplyAPIRequest) Reset() {
	r._sellerId = ""
	r._spreadId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponApplyAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionCouponApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerId is SellerId Setter
// 卖家id
func (r *TaobaoPromotionCouponApplyAPIRequest) SetSellerId(_sellerId string) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaoPromotionCouponApplyAPIRequest) GetSellerId() string {
	return r._sellerId
}

// SetSpreadId is SpreadId Setter
// 传播id
func (r *TaobaoPromotionCouponApplyAPIRequest) SetSpreadId(_spreadId string) error {
	r._spreadId = _spreadId
	r.Set("spread_id", _spreadId)
	return nil
}

// GetSpreadId SpreadId Getter
func (r TaobaoPromotionCouponApplyAPIRequest) GetSpreadId() string {
	return r._spreadId
}

var poolTaobaoPromotionCouponApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionCouponApplyRequest()
	},
}

// GetTaobaoPromotionCouponApplyRequest 从 sync.Pool 获取 TaobaoPromotionCouponApplyAPIRequest
func GetTaobaoPromotionCouponApplyAPIRequest() *TaobaoPromotionCouponApplyAPIRequest {
	return poolTaobaoPromotionCouponApplyAPIRequest.Get().(*TaobaoPromotionCouponApplyAPIRequest)
}

// ReleaseTaobaoPromotionCouponApplyAPIRequest 将 TaobaoPromotionCouponApplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionCouponApplyAPIRequest(v *TaobaoPromotionCouponApplyAPIRequest) {
	v.Reset()
	poolTaobaoPromotionCouponApplyAPIRequest.Put(v)
}
