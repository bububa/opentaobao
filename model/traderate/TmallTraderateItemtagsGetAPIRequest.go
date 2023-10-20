package traderate

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraderateItemtagsGetAPIRequest 通过商品ID获取标签列表 API请求
// tmall.traderate.itemtags.get
//
// 通过商品ID获取标签详细信息
type TmallTraderateItemtagsGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallTraderateItemtagsGetRequest 初始化TmallTraderateItemtagsGetAPIRequest对象
func NewTmallTraderateItemtagsGetRequest() *TmallTraderateItemtagsGetAPIRequest {
	return &TmallTraderateItemtagsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTraderateItemtagsGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraderateItemtagsGetAPIRequest) GetApiMethodName() string {
	return "tmall.traderate.itemtags.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraderateItemtagsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTraderateItemtagsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallTraderateItemtagsGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallTraderateItemtagsGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallTraderateItemtagsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTraderateItemtagsGetRequest()
	},
}

// GetTmallTraderateItemtagsGetRequest 从 sync.Pool 获取 TmallTraderateItemtagsGetAPIRequest
func GetTmallTraderateItemtagsGetAPIRequest() *TmallTraderateItemtagsGetAPIRequest {
	return poolTmallTraderateItemtagsGetAPIRequest.Get().(*TmallTraderateItemtagsGetAPIRequest)
}

// ReleaseTmallTraderateItemtagsGetAPIRequest 将 TmallTraderateItemtagsGetAPIRequest 放入 sync.Pool
func ReleaseTmallTraderateItemtagsGetAPIRequest(v *TmallTraderateItemtagsGetAPIRequest) {
	v.Reset()
	poolTmallTraderateItemtagsGetAPIRequest.Put(v)
}
