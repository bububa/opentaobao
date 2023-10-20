package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreItemFromOnlineItemAPIRequest 基于新模型商品id查询摊位子品id API请求
// tmall.nrt.store.item.from.online.item
//
// 基于新模型商品id查询摊位子品id
type TmallNrtStoreItemFromOnlineItemAPIRequest struct {
	model.Params
	// 线上商品编码
	_mainItemId int64
}

// NewTmallNrtStoreItemFromOnlineItemRequest 初始化TmallNrtStoreItemFromOnlineItemAPIRequest对象
func NewTmallNrtStoreItemFromOnlineItemRequest() *TmallNrtStoreItemFromOnlineItemAPIRequest {
	return &TmallNrtStoreItemFromOnlineItemAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtStoreItemFromOnlineItemAPIRequest) Reset() {
	r._mainItemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtStoreItemFromOnlineItemAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.store.item.from.online.item"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtStoreItemFromOnlineItemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtStoreItemFromOnlineItemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainItemId is MainItemId Setter
// 线上商品编码
func (r *TmallNrtStoreItemFromOnlineItemAPIRequest) SetMainItemId(_mainItemId int64) error {
	r._mainItemId = _mainItemId
	r.Set("main_item_id", _mainItemId)
	return nil
}

// GetMainItemId MainItemId Getter
func (r TmallNrtStoreItemFromOnlineItemAPIRequest) GetMainItemId() int64 {
	return r._mainItemId
}

var poolTmallNrtStoreItemFromOnlineItemAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtStoreItemFromOnlineItemRequest()
	},
}

// GetTmallNrtStoreItemFromOnlineItemRequest 从 sync.Pool 获取 TmallNrtStoreItemFromOnlineItemAPIRequest
func GetTmallNrtStoreItemFromOnlineItemAPIRequest() *TmallNrtStoreItemFromOnlineItemAPIRequest {
	return poolTmallNrtStoreItemFromOnlineItemAPIRequest.Get().(*TmallNrtStoreItemFromOnlineItemAPIRequest)
}

// ReleaseTmallNrtStoreItemFromOnlineItemAPIRequest 将 TmallNrtStoreItemFromOnlineItemAPIRequest 放入 sync.Pool
func ReleaseTmallNrtStoreItemFromOnlineItemAPIRequest(v *TmallNrtStoreItemFromOnlineItemAPIRequest) {
	v.Reset()
	poolTmallNrtStoreItemFromOnlineItemAPIRequest.Put(v)
}
