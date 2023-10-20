package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleitemuserpublishitemsAPIRequest 发布的商品列表 API请求
// alibaba.idle.item.user.publishitems
//
// 为服务商的卖家提供发布的闲鱼商品列表
type AlibabaidleitemuserpublishitemsAPIRequest struct {
	model.Params
	// 查询参数
	_paramItemPageQuery *ItemPageQuery
}

// NewAlibabaidleitemuserpublishitemsRequest 初始化AlibabaidleitemuserpublishitemsAPIRequest对象
func NewAlibabaidleitemuserpublishitemsRequest() *AlibabaidleitemuserpublishitemsAPIRequest {
	return &AlibabaidleitemuserpublishitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleitemuserpublishitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.item.user.publishitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleitemuserpublishitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleitemuserpublishitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamItemPageQuery is ParamItemPageQuery Setter
// 查询参数
func (r *AlibabaidleitemuserpublishitemsAPIRequest) SetParamItemPageQuery(_paramItemPageQuery *ItemPageQuery) error {
	r._paramItemPageQuery = _paramItemPageQuery
	r.Set("param_item_page_query", _paramItemPageQuery)
	return nil
}

// GetParamItemPageQuery ParamItemPageQuery Getter
func (r AlibabaidleitemuserpublishitemsAPIRequest) GetParamItemPageQuery() *ItemPageQuery {
	return r._paramItemPageQuery
}
