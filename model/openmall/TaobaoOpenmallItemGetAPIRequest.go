package openmall

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallItemGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallItemGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
