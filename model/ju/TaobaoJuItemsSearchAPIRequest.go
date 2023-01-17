package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJuItemsSearchAPIRequest 聚划算商品搜索接口 API请求
// taobao.ju.items.search
//
// 搜索聚划算商品
type TaobaoJuItemsSearchAPIRequest struct {
	model.Params
	// query
	_paramTopItemQuery *TopItemQuery
}

// NewTaobaoJuItemsSearchRequest 初始化TaobaoJuItemsSearchAPIRequest对象
func NewTaobaoJuItemsSearchRequest() *TaobaoJuItemsSearchAPIRequest {
	return &TaobaoJuItemsSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJuItemsSearchAPIRequest) GetApiMethodName() string {
	return "taobao.ju.items.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJuItemsSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJuItemsSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopItemQuery is ParamTopItemQuery Setter
// query
func (r *TaobaoJuItemsSearchAPIRequest) SetParamTopItemQuery(_paramTopItemQuery *TopItemQuery) error {
	r._paramTopItemQuery = _paramTopItemQuery
	r.Set("param_top_item_query", _paramTopItemQuery)
	return nil
}

// GetParamTopItemQuery ParamTopItemQuery Getter
func (r TaobaoJuItemsSearchAPIRequest) GetParamTopItemQuery() *TopItemQuery {
	return r._paramTopItemQuery
}
