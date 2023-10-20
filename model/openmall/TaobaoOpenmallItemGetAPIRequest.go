package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallItemGetAPIRequest 获取商品详情物料 API请求
// taobao.openmall.item.get
//
// 获取联盟开放的openmall商品
type TaobaoOpenmallItemGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTaobaoOpenmallItemGetRequest 初始化TaobaoOpenmallItemGetAPIRequest对象
func NewTaobaoOpenmallItemGetRequest() *TaobaoOpenmallItemGetAPIRequest {
	return &TaobaoOpenmallItemGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallItemGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallItemGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallItemGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallItemGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoOpenmallItemGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoOpenmallItemGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoOpenmallItemGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallItemGetRequest()
	},
}

// GetTaobaoOpenmallItemGetRequest 从 sync.Pool 获取 TaobaoOpenmallItemGetAPIRequest
func GetTaobaoOpenmallItemGetAPIRequest() *TaobaoOpenmallItemGetAPIRequest {
	return poolTaobaoOpenmallItemGetAPIRequest.Get().(*TaobaoOpenmallItemGetAPIRequest)
}

// ReleaseTaobaoOpenmallItemGetAPIRequest 将 TaobaoOpenmallItemGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallItemGetAPIRequest(v *TaobaoOpenmallItemGetAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallItemGetAPIRequest.Put(v)
}
