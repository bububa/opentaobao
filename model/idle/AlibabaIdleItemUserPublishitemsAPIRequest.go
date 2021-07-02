package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleItemUserPublishitemsAPIRequest 发布的商品列表 API请求
// alibaba.idle.item.user.publishitems
//
// 为服务商的卖家提供发布的闲鱼商品列表
type AlibabaIdleItemUserPublishitemsAPIRequest struct {
	model.Params
	// 查询参数
	_paramItemPageQuery *ItemPageQuery
}

// NewAlibabaIdleItemUserPublishitemsRequest 初始化AlibabaIdleItemUserPublishitemsAPIRequest对象
func NewAlibabaIdleItemUserPublishitemsRequest() *AlibabaIdleItemUserPublishitemsAPIRequest {
	return &AlibabaIdleItemUserPublishitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleItemUserPublishitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.item.user.publishitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleItemUserPublishitemsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamItemPageQuery is ParamItemPageQuery Setter
// 查询参数
func (r *AlibabaIdleItemUserPublishitemsAPIRequest) SetParamItemPageQuery(_paramItemPageQuery *ItemPageQuery) error {
	r._paramItemPageQuery = _paramItemPageQuery
	r.Set("param_item_page_query", _paramItemPageQuery)
	return nil
}

// GetParamItemPageQuery ParamItemPageQuery Getter
func (r AlibabaIdleItemUserPublishitemsAPIRequest) GetParamItemPageQuery() *ItemPageQuery {
	return r._paramItemPageQuery
}
