package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemPromotionRuleGetAPIRequest 获取商品已生效营销活动更新规则 API请求
// taobao.item.promotion.rule.get
//
// 获取商品已生效的更新规则信息，主要包含库存禁止修改，商品一口价禁止修改，库存减少锁定等规则生效信息
type TaobaoItemPromotionRuleGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTaobaoItemPromotionRuleGetRequest 初始化TaobaoItemPromotionRuleGetAPIRequest对象
func NewTaobaoItemPromotionRuleGetRequest() *TaobaoItemPromotionRuleGetAPIRequest {
	return &TaobaoItemPromotionRuleGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemPromotionRuleGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemPromotionRuleGetAPIRequest) GetApiMethodName() string {
	return "taobao.item.promotion.rule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemPromotionRuleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemPromotionRuleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoItemPromotionRuleGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoItemPromotionRuleGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoItemPromotionRuleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemPromotionRuleGetRequest()
	},
}

// GetTaobaoItemPromotionRuleGetRequest 从 sync.Pool 获取 TaobaoItemPromotionRuleGetAPIRequest
func GetTaobaoItemPromotionRuleGetAPIRequest() *TaobaoItemPromotionRuleGetAPIRequest {
	return poolTaobaoItemPromotionRuleGetAPIRequest.Get().(*TaobaoItemPromotionRuleGetAPIRequest)
}

// ReleaseTaobaoItemPromotionRuleGetAPIRequest 将 TaobaoItemPromotionRuleGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemPromotionRuleGetAPIRequest(v *TaobaoItemPromotionRuleGetAPIRequest) {
	v.Reset()
	poolTaobaoItemPromotionRuleGetAPIRequest.Put(v)
}
