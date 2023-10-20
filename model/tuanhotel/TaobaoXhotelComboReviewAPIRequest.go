package tuanhotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelComboReviewAPIRequest 套餐审核接口 API请求
// taobao.xhotel.combo.review
//
// 套餐审核接口
type TaobaoXhotelComboReviewAPIRequest struct {
	model.Params
	// 宝贝id
	_itemId int64
}

// NewTaobaoXhotelComboReviewRequest 初始化TaobaoXhotelComboReviewAPIRequest对象
func NewTaobaoXhotelComboReviewRequest() *TaobaoXhotelComboReviewAPIRequest {
	return &TaobaoXhotelComboReviewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelComboReviewAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelComboReviewAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.combo.review"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelComboReviewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelComboReviewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 宝贝id
func (r *TaobaoXhotelComboReviewAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoXhotelComboReviewAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoXhotelComboReviewAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelComboReviewRequest()
	},
}

// GetTaobaoXhotelComboReviewRequest 从 sync.Pool 获取 TaobaoXhotelComboReviewAPIRequest
func GetTaobaoXhotelComboReviewAPIRequest() *TaobaoXhotelComboReviewAPIRequest {
	return poolTaobaoXhotelComboReviewAPIRequest.Get().(*TaobaoXhotelComboReviewAPIRequest)
}

// ReleaseTaobaoXhotelComboReviewAPIRequest 将 TaobaoXhotelComboReviewAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelComboReviewAPIRequest(v *TaobaoXhotelComboReviewAPIRequest) {
	v.Reset()
	poolTaobaoXhotelComboReviewAPIRequest.Put(v)
}
