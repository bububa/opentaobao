package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest 委托信息查询接口 API请求
// alibaba.alihouse.existinghome.entrustselling.query
//
// 管家状态及房源信息接口
type AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest struct {
	model.Params
	// 入参
	_entrustSellingQuery *EntrustSellingQuery
}

// NewAlibabaalihouseexistinghomeentrustsellingqueryRequest 初始化AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest对象
func NewAlibabaalihouseexistinghomeentrustsellingqueryRequest() *AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest {
	return &AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.entrustselling.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntrustSellingQuery is EntrustSellingQuery Setter
// 入参
func (r *AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest) SetEntrustSellingQuery(_entrustSellingQuery *EntrustSellingQuery) error {
	r._entrustSellingQuery = _entrustSellingQuery
	r.Set("entrust_selling_query", _entrustSellingQuery)
	return nil
}

// GetEntrustSellingQuery EntrustSellingQuery Getter
func (r AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest) GetEntrustSellingQuery() *EntrustSellingQuery {
	return r._entrustSellingQuery
}
