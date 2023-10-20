package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemSubscribeRelationQueryAPIRequest 查询单个订阅关系 API请求
// taobao.baichuan.item.subscribe.relation.query
//
// 查询单个订阅关系
type TaobaoBaichuanItemSubscribeRelationQueryAPIRequest struct {
	model.Params
	// 商品Id
	_itemId int64
}

// NewTaobaoBaichuanItemSubscribeRelationQueryRequest 初始化TaobaoBaichuanItemSubscribeRelationQueryAPIRequest对象
func NewTaobaoBaichuanItemSubscribeRelationQueryRequest() *TaobaoBaichuanItemSubscribeRelationQueryAPIRequest {
	return &TaobaoBaichuanItemSubscribeRelationQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.item.subscribe.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品Id
func (r *TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoBaichuanItemSubscribeRelationQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanItemSubscribeRelationQueryRequest()
	},
}

// GetTaobaoBaichuanItemSubscribeRelationQueryRequest 从 sync.Pool 获取 TaobaoBaichuanItemSubscribeRelationQueryAPIRequest
func GetTaobaoBaichuanItemSubscribeRelationQueryAPIRequest() *TaobaoBaichuanItemSubscribeRelationQueryAPIRequest {
	return poolTaobaoBaichuanItemSubscribeRelationQueryAPIRequest.Get().(*TaobaoBaichuanItemSubscribeRelationQueryAPIRequest)
}

// ReleaseTaobaoBaichuanItemSubscribeRelationQueryAPIRequest 将 TaobaoBaichuanItemSubscribeRelationQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanItemSubscribeRelationQueryAPIRequest(v *TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanItemSubscribeRelationQueryAPIRequest.Put(v)
}
