package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenitemstagqueryAPIRequest 打标结果查询-商品维度 API请求
// taobao.qimen.items.tag.query
//
// 调用该接口，查询某个/某批商品上的标
type TaobaoqimenitemstagqueryAPIRequest struct {
	model.Params
	// 线上淘宝商品ID，long，必填
	_itemIds []string
}

// NewTaobaoqimenitemstagqueryRequest 初始化TaobaoqimenitemstagqueryAPIRequest对象
func NewTaobaoqimenitemstagqueryRequest() *TaobaoqimenitemstagqueryAPIRequest {
	return &TaobaoqimenitemstagqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenitemstagqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.items.tag.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenitemstagqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenitemstagqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 线上淘宝商品ID，long，必填
func (r *TaobaoqimenitemstagqueryAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoqimenitemstagqueryAPIRequest) GetItemIds() []string {
	return r._itemIds
}
