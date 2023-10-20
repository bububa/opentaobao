package opentrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialRuleUpdateAPIRequest 专属下单更新限购规则 API请求
// taobao.opentrade.special.rule.update
//
// 对于专属下单的交易场景更新限购规则
type TaobaoOpentradeSpecialRuleUpdateAPIRequest struct {
	model.Params
	// 商品id列表
	_itemIds []int64
	// 最大限购数量
	_limitNum int64
}

// NewTaobaoOpentradeSpecialRuleUpdateRequest 初始化TaobaoOpentradeSpecialRuleUpdateAPIRequest对象
func NewTaobaoOpentradeSpecialRuleUpdateRequest() *TaobaoOpentradeSpecialRuleUpdateAPIRequest {
	return &TaobaoOpentradeSpecialRuleUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpentradeSpecialRuleUpdateAPIRequest) Reset() {
	r._itemIds = r._itemIds[:0]
	r._limitNum = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialRuleUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.special.rule.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialRuleUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeSpecialRuleUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 商品id列表
func (r *TaobaoOpentradeSpecialRuleUpdateAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoOpentradeSpecialRuleUpdateAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetLimitNum is LimitNum Setter
// 最大限购数量
func (r *TaobaoOpentradeSpecialRuleUpdateAPIRequest) SetLimitNum(_limitNum int64) error {
	r._limitNum = _limitNum
	r.Set("limit_num", _limitNum)
	return nil
}

// GetLimitNum LimitNum Getter
func (r TaobaoOpentradeSpecialRuleUpdateAPIRequest) GetLimitNum() int64 {
	return r._limitNum
}

var poolTaobaoOpentradeSpecialRuleUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpentradeSpecialRuleUpdateRequest()
	},
}

// GetTaobaoOpentradeSpecialRuleUpdateRequest 从 sync.Pool 获取 TaobaoOpentradeSpecialRuleUpdateAPIRequest
func GetTaobaoOpentradeSpecialRuleUpdateAPIRequest() *TaobaoOpentradeSpecialRuleUpdateAPIRequest {
	return poolTaobaoOpentradeSpecialRuleUpdateAPIRequest.Get().(*TaobaoOpentradeSpecialRuleUpdateAPIRequest)
}

// ReleaseTaobaoOpentradeSpecialRuleUpdateAPIRequest 将 TaobaoOpentradeSpecialRuleUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpentradeSpecialRuleUpdateAPIRequest(v *TaobaoOpentradeSpecialRuleUpdateAPIRequest) {
	v.Reset()
	poolTaobaoOpentradeSpecialRuleUpdateAPIRequest.Put(v)
}
