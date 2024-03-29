package paimai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPaimaiItemCooperationSyncAPIRequest 商品同步 API请求
// taobao.paimai.item.cooperation.sync
//
// 商品同步
type TaobaoPaimaiItemCooperationSyncAPIRequest struct {
	model.Params
	// 日期
	_ds int64
	// 无
	_itemDo *ItemDo
}

// NewTaobaoPaimaiItemCooperationSyncRequest 初始化TaobaoPaimaiItemCooperationSyncAPIRequest对象
func NewTaobaoPaimaiItemCooperationSyncRequest() *TaobaoPaimaiItemCooperationSyncAPIRequest {
	return &TaobaoPaimaiItemCooperationSyncAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPaimaiItemCooperationSyncAPIRequest) Reset() {
	r._ds = 0
	r._itemDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPaimaiItemCooperationSyncAPIRequest) GetApiMethodName() string {
	return "taobao.paimai.item.cooperation.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPaimaiItemCooperationSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPaimaiItemCooperationSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDs is Ds Setter
// 日期
func (r *TaobaoPaimaiItemCooperationSyncAPIRequest) SetDs(_ds int64) error {
	r._ds = _ds
	r.Set("ds", _ds)
	return nil
}

// GetDs Ds Getter
func (r TaobaoPaimaiItemCooperationSyncAPIRequest) GetDs() int64 {
	return r._ds
}

// SetItemDo is ItemDo Setter
// 无
func (r *TaobaoPaimaiItemCooperationSyncAPIRequest) SetItemDo(_itemDo *ItemDo) error {
	r._itemDo = _itemDo
	r.Set("item_do", _itemDo)
	return nil
}

// GetItemDo ItemDo Getter
func (r TaobaoPaimaiItemCooperationSyncAPIRequest) GetItemDo() *ItemDo {
	return r._itemDo
}

var poolTaobaoPaimaiItemCooperationSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPaimaiItemCooperationSyncRequest()
	},
}

// GetTaobaoPaimaiItemCooperationSyncRequest 从 sync.Pool 获取 TaobaoPaimaiItemCooperationSyncAPIRequest
func GetTaobaoPaimaiItemCooperationSyncAPIRequest() *TaobaoPaimaiItemCooperationSyncAPIRequest {
	return poolTaobaoPaimaiItemCooperationSyncAPIRequest.Get().(*TaobaoPaimaiItemCooperationSyncAPIRequest)
}

// ReleaseTaobaoPaimaiItemCooperationSyncAPIRequest 将 TaobaoPaimaiItemCooperationSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoPaimaiItemCooperationSyncAPIRequest(v *TaobaoPaimaiItemCooperationSyncAPIRequest) {
	v.Reset()
	poolTaobaoPaimaiItemCooperationSyncAPIRequest.Put(v)
}
