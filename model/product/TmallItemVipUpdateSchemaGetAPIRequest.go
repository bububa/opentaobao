package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemVipUpdateSchemaGetAPIRequest vip商家编辑商品的规则获取接口 API请求
// tmall.item.vip.update.schema.get
//
// 获取vip商家编辑商品的规则
type TmallItemVipUpdateSchemaGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallItemVipUpdateSchemaGetRequest 初始化TmallItemVipUpdateSchemaGetAPIRequest对象
func NewTmallItemVipUpdateSchemaGetRequest() *TmallItemVipUpdateSchemaGetAPIRequest {
	return &TmallItemVipUpdateSchemaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemVipUpdateSchemaGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemVipUpdateSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemVipUpdateSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemVipUpdateSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallItemVipUpdateSchemaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemVipUpdateSchemaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallItemVipUpdateSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemVipUpdateSchemaGetRequest()
	},
}

// GetTmallItemVipUpdateSchemaGetRequest 从 sync.Pool 获取 TmallItemVipUpdateSchemaGetAPIRequest
func GetTmallItemVipUpdateSchemaGetAPIRequest() *TmallItemVipUpdateSchemaGetAPIRequest {
	return poolTmallItemVipUpdateSchemaGetAPIRequest.Get().(*TmallItemVipUpdateSchemaGetAPIRequest)
}

// ReleaseTmallItemVipUpdateSchemaGetAPIRequest 将 TmallItemVipUpdateSchemaGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemVipUpdateSchemaGetAPIRequest(v *TmallItemVipUpdateSchemaGetAPIRequest) {
	v.Reset()
	poolTmallItemVipUpdateSchemaGetAPIRequest.Put(v)
}
