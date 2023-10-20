package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentItemQueryAPIRequest 查询租赁商品信息 API请求
// alibaba.idle.rent.item.query
//
// 查询租赁商品信息
type AlibabaIdleRentItemQueryAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewAlibabaIdleRentItemQueryRequest 初始化AlibabaIdleRentItemQueryAPIRequest对象
func NewAlibabaIdleRentItemQueryRequest() *AlibabaIdleRentItemQueryAPIRequest {
	return &AlibabaIdleRentItemQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRentItemQueryAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlibabaIdleRentItemQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaIdleRentItemQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolAlibabaIdleRentItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRentItemQueryRequest()
	},
}

// GetAlibabaIdleRentItemQueryRequest 从 sync.Pool 获取 AlibabaIdleRentItemQueryAPIRequest
func GetAlibabaIdleRentItemQueryAPIRequest() *AlibabaIdleRentItemQueryAPIRequest {
	return poolAlibabaIdleRentItemQueryAPIRequest.Get().(*AlibabaIdleRentItemQueryAPIRequest)
}

// ReleaseAlibabaIdleRentItemQueryAPIRequest 将 AlibabaIdleRentItemQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRentItemQueryAPIRequest(v *AlibabaIdleRentItemQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleRentItemQueryAPIRequest.Put(v)
}
