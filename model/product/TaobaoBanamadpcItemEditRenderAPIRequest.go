package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBanamadpcItemEditRenderAPIRequest 编辑商品发布页 API请求
// taobao.banamadpc.item.edit.render
//
// 巴拿马供应商通过此接口获取编辑商品发布页
type TaobaoBanamadpcItemEditRenderAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTaobaoBanamadpcItemEditRenderRequest 初始化TaobaoBanamadpcItemEditRenderAPIRequest对象
func NewTaobaoBanamadpcItemEditRenderRequest() *TaobaoBanamadpcItemEditRenderAPIRequest {
	return &TaobaoBanamadpcItemEditRenderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBanamadpcItemEditRenderAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemEditRenderAPIRequest) GetApiMethodName() string {
	return "taobao.banamadpc.item.edit.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemEditRenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBanamadpcItemEditRenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoBanamadpcItemEditRenderAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoBanamadpcItemEditRenderAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoBanamadpcItemEditRenderAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBanamadpcItemEditRenderRequest()
	},
}

// GetTaobaoBanamadpcItemEditRenderRequest 从 sync.Pool 获取 TaobaoBanamadpcItemEditRenderAPIRequest
func GetTaobaoBanamadpcItemEditRenderAPIRequest() *TaobaoBanamadpcItemEditRenderAPIRequest {
	return poolTaobaoBanamadpcItemEditRenderAPIRequest.Get().(*TaobaoBanamadpcItemEditRenderAPIRequest)
}

// ReleaseTaobaoBanamadpcItemEditRenderAPIRequest 将 TaobaoBanamadpcItemEditRenderAPIRequest 放入 sync.Pool
func ReleaseTaobaoBanamadpcItemEditRenderAPIRequest(v *TaobaoBanamadpcItemEditRenderAPIRequest) {
	v.Reset()
	poolTaobaoBanamadpcItemEditRenderAPIRequest.Put(v)
}
