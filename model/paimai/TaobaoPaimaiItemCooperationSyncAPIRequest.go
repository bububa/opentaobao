package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopaimaiitemcooperationsyncAPIRequest 商品同步 API请求
// taobao.paimai.item.cooperation.sync
//
// 商品同步
type TaobaopaimaiitemcooperationsyncAPIRequest struct {
	model.Params
	// 日期
	_ds int64
	// 无
	_itemDo *ItemDo
}

// NewTaobaopaimaiitemcooperationsyncRequest 初始化TaobaopaimaiitemcooperationsyncAPIRequest对象
func NewTaobaopaimaiitemcooperationsyncRequest() *TaobaopaimaiitemcooperationsyncAPIRequest {
	return &TaobaopaimaiitemcooperationsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopaimaiitemcooperationsyncAPIRequest) GetApiMethodName() string {
	return "taobao.paimai.item.cooperation.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopaimaiitemcooperationsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopaimaiitemcooperationsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDs is Ds Setter
// 日期
func (r *TaobaopaimaiitemcooperationsyncAPIRequest) SetDs(_ds int64) error {
	r._ds = _ds
	r.Set("ds", _ds)
	return nil
}

// GetDs Ds Getter
func (r TaobaopaimaiitemcooperationsyncAPIRequest) GetDs() int64 {
	return r._ds
}

// SetItemDo is ItemDo Setter
// 无
func (r *TaobaopaimaiitemcooperationsyncAPIRequest) SetItemDo(_itemDo *ItemDo) error {
	r._itemDo = _itemDo
	r.Set("item_do", _itemDo)
	return nil
}

// GetItemDo ItemDo Getter
func (r TaobaopaimaiitemcooperationsyncAPIRequest) GetItemDo() *ItemDo {
	return r._itemDo
}
