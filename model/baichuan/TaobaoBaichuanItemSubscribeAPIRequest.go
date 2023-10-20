package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemSubscribeAPIRequest 单个商品订阅 API请求
// taobao.baichuan.item.subscribe
//
// 百川单个商品订阅
type TaobaoBaichuanItemSubscribeAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTaobaoBaichuanItemSubscribeRequest 初始化TaobaoBaichuanItemSubscribeAPIRequest对象
func NewTaobaoBaichuanItemSubscribeRequest() *TaobaoBaichuanItemSubscribeAPIRequest {
	return &TaobaoBaichuanItemSubscribeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanItemSubscribeAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemSubscribeAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.item.subscribe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemSubscribeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanItemSubscribeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoBaichuanItemSubscribeAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoBaichuanItemSubscribeAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoBaichuanItemSubscribeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanItemSubscribeRequest()
	},
}

// GetTaobaoBaichuanItemSubscribeRequest 从 sync.Pool 获取 TaobaoBaichuanItemSubscribeAPIRequest
func GetTaobaoBaichuanItemSubscribeAPIRequest() *TaobaoBaichuanItemSubscribeAPIRequest {
	return poolTaobaoBaichuanItemSubscribeAPIRequest.Get().(*TaobaoBaichuanItemSubscribeAPIRequest)
}

// ReleaseTaobaoBaichuanItemSubscribeAPIRequest 将 TaobaoBaichuanItemSubscribeAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanItemSubscribeAPIRequest(v *TaobaoBaichuanItemSubscribeAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanItemSubscribeAPIRequest.Put(v)
}
