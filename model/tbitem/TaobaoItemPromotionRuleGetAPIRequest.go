package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitempromotionrulegetAPIRequest 获取商品已生效营销活动更新规则 API请求
// taobao.item.promotion.rule.get
//
// 获取商品已生效的更新规则信息，主要包含库存禁止修改，商品一口价禁止修改，库存减少锁定等规则生效信息
type TaobaoitempromotionrulegetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTaobaoitempromotionrulegetRequest 初始化TaobaoitempromotionrulegetAPIRequest对象
func NewTaobaoitempromotionrulegetRequest() *TaobaoitempromotionrulegetAPIRequest {
	return &TaobaoitempromotionrulegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitempromotionrulegetAPIRequest) GetApiMethodName() string {
	return "taobao.item.promotion.rule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitempromotionrulegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitempromotionrulegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoitempromotionrulegetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoitempromotionrulegetAPIRequest) GetItemId() int64 {
	return r._itemId
}
