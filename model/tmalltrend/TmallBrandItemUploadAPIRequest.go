package tmalltrend

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallBrandItemUploadAPIRequest 天猫品牌新品同步API API请求
// tmall.brand.item.upload
//
// 支撑天猫品牌将各渠道新品信息同步至平台
type TmallBrandItemUploadAPIRequest struct {
	model.Params
	// 需要同步的商品列表
	_itemList []TmallBrandChannelNewItem
}

// NewTmallBrandItemUploadRequest 初始化TmallBrandItemUploadAPIRequest对象
func NewTmallBrandItemUploadRequest() *TmallBrandItemUploadAPIRequest {
	return &TmallBrandItemUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallBrandItemUploadAPIRequest) Reset() {
	r._itemList = r._itemList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallBrandItemUploadAPIRequest) GetApiMethodName() string {
	return "tmall.brand.item.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallBrandItemUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallBrandItemUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// 需要同步的商品列表
func (r *TmallBrandItemUploadAPIRequest) SetItemList(_itemList []TmallBrandChannelNewItem) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r TmallBrandItemUploadAPIRequest) GetItemList() []TmallBrandChannelNewItem {
	return r._itemList
}

var poolTmallBrandItemUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallBrandItemUploadRequest()
	},
}

// GetTmallBrandItemUploadRequest 从 sync.Pool 获取 TmallBrandItemUploadAPIRequest
func GetTmallBrandItemUploadAPIRequest() *TmallBrandItemUploadAPIRequest {
	return poolTmallBrandItemUploadAPIRequest.Get().(*TmallBrandItemUploadAPIRequest)
}

// ReleaseTmallBrandItemUploadAPIRequest 将 TmallBrandItemUploadAPIRequest 放入 sync.Pool
func ReleaseTmallBrandItemUploadAPIRequest(v *TmallBrandItemUploadAPIRequest) {
	v.Reset()
	poolTmallBrandItemUploadAPIRequest.Put(v)
}
