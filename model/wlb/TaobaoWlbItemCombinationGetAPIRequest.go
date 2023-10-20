package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemCombinationGetAPIRequest 根据商品id查询商品组合关系 API请求
// taobao.wlb.item.combination.get
//
// 根据商品id查询商品组合关系
type TaobaoWlbItemCombinationGetAPIRequest struct {
	model.Params
	// 要查询的组合商品id
	_itemId int64
}

// NewTaobaoWlbItemCombinationGetRequest 初始化TaobaoWlbItemCombinationGetAPIRequest对象
func NewTaobaoWlbItemCombinationGetRequest() *TaobaoWlbItemCombinationGetAPIRequest {
	return &TaobaoWlbItemCombinationGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbItemCombinationGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemCombinationGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.combination.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemCombinationGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbItemCombinationGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 要查询的组合商品id
func (r *TaobaoWlbItemCombinationGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoWlbItemCombinationGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoWlbItemCombinationGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbItemCombinationGetRequest()
	},
}

// GetTaobaoWlbItemCombinationGetRequest 从 sync.Pool 获取 TaobaoWlbItemCombinationGetAPIRequest
func GetTaobaoWlbItemCombinationGetAPIRequest() *TaobaoWlbItemCombinationGetAPIRequest {
	return poolTaobaoWlbItemCombinationGetAPIRequest.Get().(*TaobaoWlbItemCombinationGetAPIRequest)
}

// ReleaseTaobaoWlbItemCombinationGetAPIRequest 将 TaobaoWlbItemCombinationGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbItemCombinationGetAPIRequest(v *TaobaoWlbItemCombinationGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbItemCombinationGetAPIRequest.Put(v)
}
