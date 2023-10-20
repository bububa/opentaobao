package ascpffo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpFfoQueryAPIRequest AliExpress发货单查询API API请求
// aliexpress.ascp.ffo.query
//
// AE 履约发货单分页查询接口
type AliexpressAscpFfoQueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto
}

// NewAliexpressAscpFfoQueryRequest 初始化AliexpressAscpFfoQueryAPIRequest对象
func NewAliexpressAscpFfoQueryRequest() *AliexpressAscpFfoQueryAPIRequest {
	return &AliexpressAscpFfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAscpFfoQueryAPIRequest) Reset() {
	r._fulfillmentForwardOrderQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpFfoQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.ffo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpFfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpFfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillmentForwardOrderQuery is FulfillmentForwardOrderQuery Setter
// dto
func (r *AliexpressAscpFfoQueryAPIRequest) SetFulfillmentForwardOrderQuery(_fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto) error {
	r._fulfillmentForwardOrderQuery = _fulfillmentForwardOrderQuery
	r.Set("fulfillment_forward_order_query", _fulfillmentForwardOrderQuery)
	return nil
}

// GetFulfillmentForwardOrderQuery FulfillmentForwardOrderQuery Getter
func (r AliexpressAscpFfoQueryAPIRequest) GetFulfillmentForwardOrderQuery() *FulfillmentForwardOrderQueryDto {
	return r._fulfillmentForwardOrderQuery
}

var poolAliexpressAscpFfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAscpFfoQueryRequest()
	},
}

// GetAliexpressAscpFfoQueryRequest 从 sync.Pool 获取 AliexpressAscpFfoQueryAPIRequest
func GetAliexpressAscpFfoQueryAPIRequest() *AliexpressAscpFfoQueryAPIRequest {
	return poolAliexpressAscpFfoQueryAPIRequest.Get().(*AliexpressAscpFfoQueryAPIRequest)
}

// ReleaseAliexpressAscpFfoQueryAPIRequest 将 AliexpressAscpFfoQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAscpFfoQueryAPIRequest(v *AliexpressAscpFfoQueryAPIRequest) {
	v.Reset()
	poolAliexpressAscpFfoQueryAPIRequest.Put(v)
}
