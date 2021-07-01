package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanItemSubscribeRelationQueryAPIRequest
查询单个订阅关系 API请求
taobao.baichuan.item.subscribe.relation.query

查询单个订阅关系 */
type TaobaoBaichuanItemSubscribeRelationQueryAPIRequest struct {
	model.Params
	// 商品Id
	_itemId int64
}

// NewTaobaoBaichuanItemSubscribeRelationQueryRequest 初始化TaobaoBaichuanItemSubscribeRelationQueryAPIRequest对象
func NewTaobaoBaichuanItemSubscribeRelationQueryRequest() *TaobaoBaichuanItemSubscribeRelationQueryAPIRequest {
	return &TaobaoBaichuanItemSubscribeRelationQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.item.subscribe.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品Id
func (r *TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoBaichuanItemSubscribeRelationQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}
