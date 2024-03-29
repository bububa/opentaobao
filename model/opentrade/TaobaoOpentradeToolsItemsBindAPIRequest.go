package opentrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeToolsItemsBindAPIRequest 交易开放商品绑定 API请求
// taobao.opentrade.tools.items.bind
//
// 交易开放商品绑定
type TaobaoOpentradeToolsItemsBindAPIRequest struct {
	model.Params
	// 待绑定商品id
	_itemIds []int64
	// 绑定交易开放场景的C端小程序ID
	_miniappId int64
}

// NewTaobaoOpentradeToolsItemsBindRequest 初始化TaobaoOpentradeToolsItemsBindAPIRequest对象
func NewTaobaoOpentradeToolsItemsBindRequest() *TaobaoOpentradeToolsItemsBindAPIRequest {
	return &TaobaoOpentradeToolsItemsBindAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpentradeToolsItemsBindAPIRequest) Reset() {
	r._itemIds = r._itemIds[:0]
	r._miniappId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeToolsItemsBindAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.tools.items.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeToolsItemsBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeToolsItemsBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 待绑定商品id
func (r *TaobaoOpentradeToolsItemsBindAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoOpentradeToolsItemsBindAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetMiniappId is MiniappId Setter
// 绑定交易开放场景的C端小程序ID
func (r *TaobaoOpentradeToolsItemsBindAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoOpentradeToolsItemsBindAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}

var poolTaobaoOpentradeToolsItemsBindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpentradeToolsItemsBindRequest()
	},
}

// GetTaobaoOpentradeToolsItemsBindRequest 从 sync.Pool 获取 TaobaoOpentradeToolsItemsBindAPIRequest
func GetTaobaoOpentradeToolsItemsBindAPIRequest() *TaobaoOpentradeToolsItemsBindAPIRequest {
	return poolTaobaoOpentradeToolsItemsBindAPIRequest.Get().(*TaobaoOpentradeToolsItemsBindAPIRequest)
}

// ReleaseTaobaoOpentradeToolsItemsBindAPIRequest 将 TaobaoOpentradeToolsItemsBindAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpentradeToolsItemsBindAPIRequest(v *TaobaoOpentradeToolsItemsBindAPIRequest) {
	v.Reset()
	poolTaobaoOpentradeToolsItemsBindAPIRequest.Put(v)
}
