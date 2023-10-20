package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponSendAPIRequest 店铺优惠券发放接口 API请求
// taobao.promotion.coupon.send
//
// 通过接口批量发放店铺优惠券（每次只能发送100张，只能发给当前授权卖家店铺的会员），发送成功则返回为空，发送失败则返回失败的买家列表和发送成功的买家和优惠券的number。注：如果所有买家都发放失败的话，is_success也为true，建议调用者根据返回的集合判断是否送入的买家都发放成功了
type TaobaoPromotionCouponSendAPIRequest struct {
	model.Params
	// 买家昵称用半角','号分割
	_buyerNick []string
	// ouid
	_ouidData []OuidData
	// openuid
	_buyerIds []string
	// 优惠券的id
	_couponId int64
}

// NewTaobaoPromotionCouponSendRequest 初始化TaobaoPromotionCouponSendAPIRequest对象
func NewTaobaoPromotionCouponSendRequest() *TaobaoPromotionCouponSendAPIRequest {
	return &TaobaoPromotionCouponSendAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionCouponSendAPIRequest) Reset() {
	r._buyerNick = r._buyerNick[:0]
	r._ouidData = r._ouidData[:0]
	r._buyerIds = r._buyerIds[:0]
	r._couponId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponSendAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.coupon.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionCouponSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称用半角&#39;,&#39;号分割
func (r *TaobaoPromotionCouponSendAPIRequest) SetBuyerNick(_buyerNick []string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoPromotionCouponSendAPIRequest) GetBuyerNick() []string {
	return r._buyerNick
}

// SetOuidData is OuidData Setter
// ouid
func (r *TaobaoPromotionCouponSendAPIRequest) SetOuidData(_ouidData []OuidData) error {
	r._ouidData = _ouidData
	r.Set("ouid_data", _ouidData)
	return nil
}

// GetOuidData OuidData Getter
func (r TaobaoPromotionCouponSendAPIRequest) GetOuidData() []OuidData {
	return r._ouidData
}

// SetBuyerIds is BuyerIds Setter
// openuid
func (r *TaobaoPromotionCouponSendAPIRequest) SetBuyerIds(_buyerIds []string) error {
	r._buyerIds = _buyerIds
	r.Set("buyer_ids", _buyerIds)
	return nil
}

// GetBuyerIds BuyerIds Getter
func (r TaobaoPromotionCouponSendAPIRequest) GetBuyerIds() []string {
	return r._buyerIds
}

// SetCouponId is CouponId Setter
// 优惠券的id
func (r *TaobaoPromotionCouponSendAPIRequest) SetCouponId(_couponId int64) error {
	r._couponId = _couponId
	r.Set("coupon_id", _couponId)
	return nil
}

// GetCouponId CouponId Getter
func (r TaobaoPromotionCouponSendAPIRequest) GetCouponId() int64 {
	return r._couponId
}

var poolTaobaoPromotionCouponSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionCouponSendRequest()
	},
}

// GetTaobaoPromotionCouponSendRequest 从 sync.Pool 获取 TaobaoPromotionCouponSendAPIRequest
func GetTaobaoPromotionCouponSendAPIRequest() *TaobaoPromotionCouponSendAPIRequest {
	return poolTaobaoPromotionCouponSendAPIRequest.Get().(*TaobaoPromotionCouponSendAPIRequest)
}

// ReleaseTaobaoPromotionCouponSendAPIRequest 将 TaobaoPromotionCouponSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionCouponSendAPIRequest(v *TaobaoPromotionCouponSendAPIRequest) {
	v.Reset()
	poolTaobaoPromotionCouponSendAPIRequest.Put(v)
}
