package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentitemqueryAPIRequest 查询租赁商品信息 API请求
// alibaba.idle.rent.item.query
//
// 查询租赁商品信息
type AlibabaidlerentitemqueryAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewAlibabaidlerentitemqueryRequest 初始化AlibabaidlerentitemqueryAPIRequest对象
func NewAlibabaidlerentitemqueryRequest() *AlibabaidlerentitemqueryAPIRequest {
	return &AlibabaidlerentitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentitemqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlibabaidlerentitemqueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaidlerentitemqueryAPIRequest) GetItemId() int64 {
	return r._itemId
}
