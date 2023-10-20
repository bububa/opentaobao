package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmallitemcentersubscribequeryAPIRequest 天猫服务订购信息查询接口 API请求
// tmall.mallitemcenter.subscribe.query
//
// 查询商家服务订购信息
type TmallmallitemcentersubscribequeryAPIRequest struct {
	model.Params
	// 入参query
	_query *Spb2bOderQuery
}

// NewTmallmallitemcentersubscribequeryRequest 初始化TmallmallitemcentersubscribequeryAPIRequest对象
func NewTmallmallitemcentersubscribequeryRequest() *TmallmallitemcentersubscribequeryAPIRequest {
	return &TmallmallitemcentersubscribequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmallitemcentersubscribequeryAPIRequest) GetApiMethodName() string {
	return "tmall.mallitemcenter.subscribe.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmallitemcentersubscribequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmallitemcentersubscribequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参query
func (r *TmallmallitemcentersubscribequeryAPIRequest) SetQuery(_query *Spb2bOderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallmallitemcentersubscribequeryAPIRequest) GetQuery() *Spb2bOderQuery {
	return r._query
}
