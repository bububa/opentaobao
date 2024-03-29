package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemGetAPIRequest 根据商品ID获取商品信息 API请求
// taobao.wlb.item.get
//
// 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
type TaobaoWlbItemGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTaobaoWlbItemGetRequest 初始化TaobaoWlbItemGetAPIRequest对象
func NewTaobaoWlbItemGetRequest() *TaobaoWlbItemGetAPIRequest {
	return &TaobaoWlbItemGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbItemGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbItemGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoWlbItemGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoWlbItemGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoWlbItemGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbItemGetRequest()
	},
}

// GetTaobaoWlbItemGetRequest 从 sync.Pool 获取 TaobaoWlbItemGetAPIRequest
func GetTaobaoWlbItemGetAPIRequest() *TaobaoWlbItemGetAPIRequest {
	return poolTaobaoWlbItemGetAPIRequest.Get().(*TaobaoWlbItemGetAPIRequest)
}

// ReleaseTaobaoWlbItemGetAPIRequest 将 TaobaoWlbItemGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbItemGetAPIRequest(v *TaobaoWlbItemGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbItemGetAPIRequest.Put(v)
}
