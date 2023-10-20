package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemItemPageAPIRequest 信息流查看商品列表 API请求
// taobao.feedflow.item.item.page
//
// 信息流查看商品列表
type TaobaoFeedflowItemItemPageAPIRequest struct {
	model.Params
	// 查询条件
	_itemQuery *ItemQueryDto
}

// NewTaobaoFeedflowItemItemPageRequest 初始化TaobaoFeedflowItemItemPageAPIRequest对象
func NewTaobaoFeedflowItemItemPageRequest() *TaobaoFeedflowItemItemPageAPIRequest {
	return &TaobaoFeedflowItemItemPageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemItemPageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.item.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemItemPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemItemPageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemQuery is ItemQuery Setter
// 查询条件
func (r *TaobaoFeedflowItemItemPageAPIRequest) SetItemQuery(_itemQuery *ItemQueryDto) error {
	r._itemQuery = _itemQuery
	r.Set("item_query", _itemQuery)
	return nil
}

// GetItemQuery ItemQuery Getter
func (r TaobaoFeedflowItemItemPageAPIRequest) GetItemQuery() *ItemQueryDto {
	return r._itemQuery
}
