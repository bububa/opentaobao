package idleisv

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleItemUserPublishitemsAPIRequest) Reset() {
	r._paramItemPageQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleItemUserPublishitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.item.user.publishitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleItemUserPublishitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleItemUserPublishitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIdleItemUserPublishitemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleItemUserPublishitemsRequest()
	},
}

// GetAlibabaIdleItemUserPublishitemsRequest 从 sync.Pool 获取 AlibabaIdleItemUserPublishitemsAPIRequest
func GetAlibabaIdleItemUserPublishitemsAPIRequest() *AlibabaIdleItemUserPublishitemsAPIRequest {
	return poolAlibabaIdleItemUserPublishitemsAPIRequest.Get().(*AlibabaIdleItemUserPublishitemsAPIRequest)
}

// ReleaseAlibabaIdleItemUserPublishitemsAPIRequest 将 AlibabaIdleItemUserPublishitemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleItemUserPublishitemsAPIRequest(v *AlibabaIdleItemUserPublishitemsAPIRequest) {
	v.Reset()
	poolAlibabaIdleItemUserPublishitemsAPIRequest.Put(v)
}
