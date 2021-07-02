package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemMapGetAPIRequest 根据物流宝商品ID查询商品映射关系 API请求
// taobao.wlb.item.map.get
//
// 根据物流宝商品ID查询商品映射关系
type TaobaoWlbItemMapGetAPIRequest struct {
	model.Params
	// 要查询映射关系的物流宝商品id
	_itemId int64
}

// NewTaobaoWlbItemMapGetRequest 初始化TaobaoWlbItemMapGetAPIRequest对象
func NewTaobaoWlbItemMapGetRequest() *TaobaoWlbItemMapGetAPIRequest {
	return &TaobaoWlbItemMapGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemMapGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.map.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemMapGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 要查询映射关系的物流宝商品id
func (r *TaobaoWlbItemMapGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoWlbItemMapGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
