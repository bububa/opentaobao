package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemsUnsubscribeAPIRequest 批量删除商品订阅 API请求
// taobao.baichuan.items.unsubscribe
//
// 批量删除商品订阅
type TaobaoBaichuanItemsUnsubscribeAPIRequest struct {
	model.Params
	// 删除的商品id
	_itemIds []string
}

// NewTaobaoBaichuanItemsUnsubscribeRequest 初始化TaobaoBaichuanItemsUnsubscribeAPIRequest对象
func NewTaobaoBaichuanItemsUnsubscribeRequest() *TaobaoBaichuanItemsUnsubscribeAPIRequest {
	return &TaobaoBaichuanItemsUnsubscribeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanItemsUnsubscribeAPIRequest) Reset() {
	r._itemIds = r._itemIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemsUnsubscribeAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.items.unsubscribe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemsUnsubscribeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanItemsUnsubscribeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 删除的商品id
func (r *TaobaoBaichuanItemsUnsubscribeAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoBaichuanItemsUnsubscribeAPIRequest) GetItemIds() []string {
	return r._itemIds
}

var poolTaobaoBaichuanItemsUnsubscribeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanItemsUnsubscribeRequest()
	},
}

// GetTaobaoBaichuanItemsUnsubscribeRequest 从 sync.Pool 获取 TaobaoBaichuanItemsUnsubscribeAPIRequest
func GetTaobaoBaichuanItemsUnsubscribeAPIRequest() *TaobaoBaichuanItemsUnsubscribeAPIRequest {
	return poolTaobaoBaichuanItemsUnsubscribeAPIRequest.Get().(*TaobaoBaichuanItemsUnsubscribeAPIRequest)
}

// ReleaseTaobaoBaichuanItemsUnsubscribeAPIRequest 将 TaobaoBaichuanItemsUnsubscribeAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanItemsUnsubscribeAPIRequest(v *TaobaoBaichuanItemsUnsubscribeAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanItemsUnsubscribeAPIRequest.Put(v)
}
