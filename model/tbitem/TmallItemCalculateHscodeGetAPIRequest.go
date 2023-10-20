package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemCalculateHscodeGetAPIRequest 算法获取hscode API请求
// tmall.item.calculate.hscode.get
//
// 算法获取hscode
type TmallItemCalculateHscodeGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallItemCalculateHscodeGetRequest 初始化TmallItemCalculateHscodeGetAPIRequest对象
func NewTmallItemCalculateHscodeGetRequest() *TmallItemCalculateHscodeGetAPIRequest {
	return &TmallItemCalculateHscodeGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemCalculateHscodeGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemCalculateHscodeGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.calculate.hscode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemCalculateHscodeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemCalculateHscodeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallItemCalculateHscodeGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemCalculateHscodeGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallItemCalculateHscodeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemCalculateHscodeGetRequest()
	},
}

// GetTmallItemCalculateHscodeGetRequest 从 sync.Pool 获取 TmallItemCalculateHscodeGetAPIRequest
func GetTmallItemCalculateHscodeGetAPIRequest() *TmallItemCalculateHscodeGetAPIRequest {
	return poolTmallItemCalculateHscodeGetAPIRequest.Get().(*TmallItemCalculateHscodeGetAPIRequest)
}

// ReleaseTmallItemCalculateHscodeGetAPIRequest 将 TmallItemCalculateHscodeGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemCalculateHscodeGetAPIRequest(v *TmallItemCalculateHscodeGetAPIRequest) {
	v.Reset()
	poolTmallItemCalculateHscodeGetAPIRequest.Put(v)
}
