package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMarketingPromotionKfcAPIRequest 定向优惠活动名称与描述违禁词检查 API请求
// taobao.marketing.promotion.kfc
//
// 活动名称与描述违禁词检查
type TaobaoMarketingPromotionKfcAPIRequest struct {
	model.Params
	// 活动名称
	_promotionTitle string
	// 活动描述
	_promotionDesc string
}

// NewTaobaoMarketingPromotionKfcRequest 初始化TaobaoMarketingPromotionKfcAPIRequest对象
func NewTaobaoMarketingPromotionKfcRequest() *TaobaoMarketingPromotionKfcAPIRequest {
	return &TaobaoMarketingPromotionKfcAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMarketingPromotionKfcAPIRequest) Reset() {
	r._promotionTitle = ""
	r._promotionDesc = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMarketingPromotionKfcAPIRequest) GetApiMethodName() string {
	return "taobao.marketing.promotion.kfc"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMarketingPromotionKfcAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMarketingPromotionKfcAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPromotionTitle is PromotionTitle Setter
// 活动名称
func (r *TaobaoMarketingPromotionKfcAPIRequest) SetPromotionTitle(_promotionTitle string) error {
	r._promotionTitle = _promotionTitle
	r.Set("promotion_title", _promotionTitle)
	return nil
}

// GetPromotionTitle PromotionTitle Getter
func (r TaobaoMarketingPromotionKfcAPIRequest) GetPromotionTitle() string {
	return r._promotionTitle
}

// SetPromotionDesc is PromotionDesc Setter
// 活动描述
func (r *TaobaoMarketingPromotionKfcAPIRequest) SetPromotionDesc(_promotionDesc string) error {
	r._promotionDesc = _promotionDesc
	r.Set("promotion_desc", _promotionDesc)
	return nil
}

// GetPromotionDesc PromotionDesc Getter
func (r TaobaoMarketingPromotionKfcAPIRequest) GetPromotionDesc() string {
	return r._promotionDesc
}

var poolTaobaoMarketingPromotionKfcAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMarketingPromotionKfcRequest()
	},
}

// GetTaobaoMarketingPromotionKfcRequest 从 sync.Pool 获取 TaobaoMarketingPromotionKfcAPIRequest
func GetTaobaoMarketingPromotionKfcAPIRequest() *TaobaoMarketingPromotionKfcAPIRequest {
	return poolTaobaoMarketingPromotionKfcAPIRequest.Get().(*TaobaoMarketingPromotionKfcAPIRequest)
}

// ReleaseTaobaoMarketingPromotionKfcAPIRequest 将 TaobaoMarketingPromotionKfcAPIRequest 放入 sync.Pool
func ReleaseTaobaoMarketingPromotionKfcAPIRequest(v *TaobaoMarketingPromotionKfcAPIRequest) {
	v.Reset()
	poolTaobaoMarketingPromotionKfcAPIRequest.Put(v)
}
