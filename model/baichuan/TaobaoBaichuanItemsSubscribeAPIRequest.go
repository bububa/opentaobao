package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemsSubscribeAPIRequest 百川批量商品订阅 API请求
// taobao.baichuan.items.subscribe
//
// 百川批量添加订阅的商品
type TaobaoBaichuanItemsSubscribeAPIRequest struct {
	model.Params
	// 订阅的商品id列表
	_itemIds []string
}

// NewTaobaoBaichuanItemsSubscribeRequest 初始化TaobaoBaichuanItemsSubscribeAPIRequest对象
func NewTaobaoBaichuanItemsSubscribeRequest() *TaobaoBaichuanItemsSubscribeAPIRequest {
	return &TaobaoBaichuanItemsSubscribeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanItemsSubscribeAPIRequest) Reset() {
	r._itemIds = r._itemIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemsSubscribeAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.items.subscribe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemsSubscribeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanItemsSubscribeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 订阅的商品id列表
func (r *TaobaoBaichuanItemsSubscribeAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoBaichuanItemsSubscribeAPIRequest) GetItemIds() []string {
	return r._itemIds
}

var poolTaobaoBaichuanItemsSubscribeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanItemsSubscribeRequest()
	},
}

// GetTaobaoBaichuanItemsSubscribeRequest 从 sync.Pool 获取 TaobaoBaichuanItemsSubscribeAPIRequest
func GetTaobaoBaichuanItemsSubscribeAPIRequest() *TaobaoBaichuanItemsSubscribeAPIRequest {
	return poolTaobaoBaichuanItemsSubscribeAPIRequest.Get().(*TaobaoBaichuanItemsSubscribeAPIRequest)
}

// ReleaseTaobaoBaichuanItemsSubscribeAPIRequest 将 TaobaoBaichuanItemsSubscribeAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanItemsSubscribeAPIRequest(v *TaobaoBaichuanItemsSubscribeAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanItemsSubscribeAPIRequest.Put(v)
}
