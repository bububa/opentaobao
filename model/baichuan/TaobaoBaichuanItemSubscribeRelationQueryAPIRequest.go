package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanitemsubscriberelationqueryAPIRequest 查询单个订阅关系 API请求
// taobao.baichuan.item.subscribe.relation.query
//
// 查询单个订阅关系
type TaobaobaichuanitemsubscriberelationqueryAPIRequest struct {
	model.Params
	// 商品Id
	_itemId int64
}

// NewTaobaobaichuanitemsubscriberelationqueryRequest 初始化TaobaobaichuanitemsubscriberelationqueryAPIRequest对象
func NewTaobaobaichuanitemsubscriberelationqueryRequest() *TaobaobaichuanitemsubscriberelationqueryAPIRequest {
	return &TaobaobaichuanitemsubscriberelationqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanitemsubscriberelationqueryAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.item.subscribe.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanitemsubscriberelationqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanitemsubscriberelationqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品Id
func (r *TaobaobaichuanitemsubscriberelationqueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaobaichuanitemsubscriberelationqueryAPIRequest) GetItemId() int64 {
	return r._itemId
}
