package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest) Reset() {
	r._entrustSellingQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.entrustselling.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeEntrustsellingQueryRequest()
	},
}

// GetAlibabaAlihouseExistinghomeEntrustsellingQueryRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest
func GetAlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest() *AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest {
	return poolAlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest.Get().(*AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest 将 AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest(v *AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest.Put(v)
}
