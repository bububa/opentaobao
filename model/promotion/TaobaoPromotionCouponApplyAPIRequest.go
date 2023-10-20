package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotioncouponapplyAPIRequest 优惠券领取 API请求
// taobao.promotion.coupon.apply
//
// 优惠券领取
type TaobaopromotioncouponapplyAPIRequest struct {
	model.Params
	// 卖家id
	_sellerId string
	// 传播id
	_spreadId string
}

// NewTaobaopromotioncouponapplyRequest 初始化TaobaopromotioncouponapplyAPIRequest对象
func NewTaobaopromotioncouponapplyRequest() *TaobaopromotioncouponapplyAPIRequest {
	return &TaobaopromotioncouponapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotioncouponapplyAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotioncouponapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotioncouponapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerId is SellerId Setter
// 卖家id
func (r *TaobaopromotioncouponapplyAPIRequest) SetSellerId(_sellerId string) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaopromotioncouponapplyAPIRequest) GetSellerId() string {
	return r._sellerId
}

// SetSpreadId is SpreadId Setter
// 传播id
func (r *TaobaopromotioncouponapplyAPIRequest) SetSpreadId(_spreadId string) error {
	r._spreadId = _spreadId
	r.Set("spread_id", _spreadId)
	return nil
}

// GetSpreadId SpreadId Getter
func (r TaobaopromotioncouponapplyAPIRequest) GetSpreadId() string {
	return r._spreadId
}
