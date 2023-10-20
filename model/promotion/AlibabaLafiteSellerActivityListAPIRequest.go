package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalafiteselleractivitylistAPIRequest 商家自运营活动列表 API请求
// alibaba.lafite.seller.activity.list
//
// 商家查询自己配置的活动列表
type AlibabalafiteselleractivitylistAPIRequest struct {
	model.Params
	// 请求入参
	_query *ActivityReadTopQuery
}

// NewAlibabalafiteselleractivitylistRequest 初始化AlibabalafiteselleractivitylistAPIRequest对象
func NewAlibabalafiteselleractivitylistRequest() *AlibabalafiteselleractivitylistAPIRequest {
	return &AlibabalafiteselleractivitylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalafiteselleractivitylistAPIRequest) GetApiMethodName() string {
	return "alibaba.lafite.seller.activity.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalafiteselleractivitylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalafiteselleractivitylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 请求入参
func (r *AlibabalafiteselleractivitylistAPIRequest) SetQuery(_query *ActivityReadTopQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabalafiteselleractivitylistAPIRequest) GetQuery() *ActivityReadTopQuery {
	return r._query
}
