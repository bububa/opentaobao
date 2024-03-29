package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemsTagQueryAPIRequest 打标结果查询-商品维度 API请求
// taobao.qimen.items.tag.query
//
// 调用该接口，查询某个/某批商品上的标
type TaobaoQimenItemsTagQueryAPIRequest struct {
	model.Params
	// 线上淘宝商品ID，long，必填
	_itemIds []string
}

// NewTaobaoQimenItemsTagQueryRequest 初始化TaobaoQimenItemsTagQueryAPIRequest对象
func NewTaobaoQimenItemsTagQueryRequest() *TaobaoQimenItemsTagQueryAPIRequest {
	return &TaobaoQimenItemsTagQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenItemsTagQueryAPIRequest) Reset() {
	r._itemIds = r._itemIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemsTagQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.items.tag.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemsTagQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenItemsTagQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 线上淘宝商品ID，long，必填
func (r *TaobaoQimenItemsTagQueryAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoQimenItemsTagQueryAPIRequest) GetItemIds() []string {
	return r._itemIds
}

var poolTaobaoQimenItemsTagQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenItemsTagQueryRequest()
	},
}

// GetTaobaoQimenItemsTagQueryRequest 从 sync.Pool 获取 TaobaoQimenItemsTagQueryAPIRequest
func GetTaobaoQimenItemsTagQueryAPIRequest() *TaobaoQimenItemsTagQueryAPIRequest {
	return poolTaobaoQimenItemsTagQueryAPIRequest.Get().(*TaobaoQimenItemsTagQueryAPIRequest)
}

// ReleaseTaobaoQimenItemsTagQueryAPIRequest 将 TaobaoQimenItemsTagQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenItemsTagQueryAPIRequest(v *TaobaoQimenItemsTagQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenItemsTagQueryAPIRequest.Put(v)
}
