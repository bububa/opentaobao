package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbItemCombinationGetAPIRequest
根据商品id查询商品组合关系 API请求
taobao.wlb.item.combination.get

根据商品id查询商品组合关系 */
type TaobaoWlbItemCombinationGetAPIRequest struct {
	model.Params
	// 要查询的组合商品id
	_itemId int64
}

// NewTaobaoWlbItemCombinationGetRequest 初始化TaobaoWlbItemCombinationGetAPIRequest对象
func NewTaobaoWlbItemCombinationGetRequest() *TaobaoWlbItemCombinationGetAPIRequest {
	return &TaobaoWlbItemCombinationGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemCombinationGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.combination.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemCombinationGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 要查询的组合商品id
func (r *TaobaoWlbItemCombinationGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoWlbItemCombinationGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
