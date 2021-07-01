package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionLimitdiscountDetailGetAPIRequest
限时打折详情查询 API请求
taobao.promotion.limitdiscount.detail.get

限时打折详情查询。查询出指定限时打折的对应商品记录信息。 */
type TaobaoPromotionLimitdiscountDetailGetAPIRequest struct {
	model.Params
	// 限时打折ID。这个针对查询唯一限时打折情况。
	_limitDiscountId int64
}

// NewTaobaoPromotionLimitdiscountDetailGetRequest 初始化TaobaoPromotionLimitdiscountDetailGetAPIRequest对象
func NewTaobaoPromotionLimitdiscountDetailGetRequest() *TaobaoPromotionLimitdiscountDetailGetAPIRequest {
	return &TaobaoPromotionLimitdiscountDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionLimitdiscountDetailGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.limitdiscount.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionLimitdiscountDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LimitDiscountId Setter
// 限时打折ID。这个针对查询唯一限时打折情况。
func (r *TaobaoPromotionLimitdiscountDetailGetAPIRequest) SetLimitDiscountId(_limitDiscountId int64) error {
	r._limitDiscountId = _limitDiscountId
	r.Set("limit_discount_id", _limitDiscountId)
	return nil
}

// Get LimitDiscountId Getter
func (r TaobaoPromotionLimitdiscountDetailGetAPIRequest) GetLimitDiscountId() int64 {
	return r._limitDiscountId
}
