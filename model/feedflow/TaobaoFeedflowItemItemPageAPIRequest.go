package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemitempageAPIRequest 信息流查看商品列表 API请求
// taobao.feedflow.item.item.page
//
// 信息流查看商品列表
type TaobaofeedflowitemitempageAPIRequest struct {
	model.Params
	// 查询条件
	_itemQuery *ItemQueryDto
}

// NewTaobaofeedflowitemitempageRequest 初始化TaobaofeedflowitemitempageAPIRequest对象
func NewTaobaofeedflowitemitempageRequest() *TaobaofeedflowitemitempageAPIRequest {
	return &TaobaofeedflowitemitempageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemitempageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.item.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemitempageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemitempageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemQuery is ItemQuery Setter
// 查询条件
func (r *TaobaofeedflowitemitempageAPIRequest) SetItemQuery(_itemQuery *ItemQueryDto) error {
	r._itemQuery = _itemQuery
	r.Set("item_query", _itemQuery)
	return nil
}

// GetItemQuery ItemQuery Getter
func (r TaobaofeedflowitemitempageAPIRequest) GetItemQuery() *ItemQueryDto {
	return r._itemQuery
}
