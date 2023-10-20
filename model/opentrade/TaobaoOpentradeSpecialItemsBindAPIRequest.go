package opentrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialItemsBindAPIRequest 专属下单场景商品绑定 API请求
// taobao.opentrade.special.items.bind
//
// 专属下单场景商品绑定
type TaobaoOpentradeSpecialItemsBindAPIRequest struct {
	model.Params
	// 本次待绑定的商品ID列表
	_itemIds []int64
	// 绑定专属下单场景的C端小程序ID
	_miniappId int64
}

// NewTaobaoOpentradeSpecialItemsBindRequest 初始化TaobaoOpentradeSpecialItemsBindAPIRequest对象
func NewTaobaoOpentradeSpecialItemsBindRequest() *TaobaoOpentradeSpecialItemsBindAPIRequest {
	return &TaobaoOpentradeSpecialItemsBindAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpentradeSpecialItemsBindAPIRequest) Reset() {
	r._itemIds = r._itemIds[:0]
	r._miniappId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.special.items.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 本次待绑定的商品ID列表
func (r *TaobaoOpentradeSpecialItemsBindAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetMiniappId is MiniappId Setter
// 绑定专属下单场景的C端小程序ID
func (r *TaobaoOpentradeSpecialItemsBindAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}

var poolTaobaoOpentradeSpecialItemsBindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpentradeSpecialItemsBindRequest()
	},
}

// GetTaobaoOpentradeSpecialItemsBindRequest 从 sync.Pool 获取 TaobaoOpentradeSpecialItemsBindAPIRequest
func GetTaobaoOpentradeSpecialItemsBindAPIRequest() *TaobaoOpentradeSpecialItemsBindAPIRequest {
	return poolTaobaoOpentradeSpecialItemsBindAPIRequest.Get().(*TaobaoOpentradeSpecialItemsBindAPIRequest)
}

// ReleaseTaobaoOpentradeSpecialItemsBindAPIRequest 将 TaobaoOpentradeSpecialItemsBindAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpentradeSpecialItemsBindAPIRequest(v *TaobaoOpentradeSpecialItemsBindAPIRequest) {
	v.Reset()
	poolTaobaoOpentradeSpecialItemsBindAPIRequest.Put(v)
}
