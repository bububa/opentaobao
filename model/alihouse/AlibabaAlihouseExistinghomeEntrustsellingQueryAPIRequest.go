package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest 委托信息查询接口 API请求
// alibaba.alihouse.existinghome.entrustselling.query
//
// 管家状态及房源信息接口
type AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest struct {
	model.Params
	// 入参
	_entrustSellingQuery *EntrustSellingQuery
}

// NewAlibabaAlihouseExistinghomeEntrustsellingQueryRequest 初始化AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest对象
func NewAlibabaAlihouseExistinghomeEntrustsellingQueryRequest() *AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest {
	return &AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.entrustselling.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEntrustSellingQuery is EntrustSellingQuery Setter
// 入参
func (r *AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest) SetEntrustSellingQuery(_entrustSellingQuery *EntrustSellingQuery) error {
	r._entrustSellingQuery = _entrustSellingQuery
	r.Set("entrust_selling_query", _entrustSellingQuery)
	return nil
}

// GetEntrustSellingQuery EntrustSellingQuery Getter
func (r AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest) GetEntrustSellingQuery() *EntrustSellingQuery {
	return r._entrustSellingQuery
}
