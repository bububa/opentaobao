package ascpffo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpFroQueryAPIRequest AliExpress销退单查询API API请求
// aliexpress.ascp.fro.query
//
// AE履约销退单查询接口
type AliexpressAscpFroQueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentReverseOrderQuery *FulfillmentReverseOrderQueryDto
}

// NewAliexpressAscpFroQueryRequest 初始化AliexpressAscpFroQueryAPIRequest对象
func NewAliexpressAscpFroQueryRequest() *AliexpressAscpFroQueryAPIRequest {
	return &AliexpressAscpFroQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAscpFroQueryAPIRequest) Reset() {
	r._fulfillmentReverseOrderQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpFroQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.fro.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpFroQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpFroQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillmentReverseOrderQuery is FulfillmentReverseOrderQuery Setter
// dto
func (r *AliexpressAscpFroQueryAPIRequest) SetFulfillmentReverseOrderQuery(_fulfillmentReverseOrderQuery *FulfillmentReverseOrderQueryDto) error {
	r._fulfillmentReverseOrderQuery = _fulfillmentReverseOrderQuery
	r.Set("fulfillment_reverse_order_query", _fulfillmentReverseOrderQuery)
	return nil
}

// GetFulfillmentReverseOrderQuery FulfillmentReverseOrderQuery Getter
func (r AliexpressAscpFroQueryAPIRequest) GetFulfillmentReverseOrderQuery() *FulfillmentReverseOrderQueryDto {
	return r._fulfillmentReverseOrderQuery
}

var poolAliexpressAscpFroQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAscpFroQueryRequest()
	},
}

// GetAliexpressAscpFroQueryRequest 从 sync.Pool 获取 AliexpressAscpFroQueryAPIRequest
func GetAliexpressAscpFroQueryAPIRequest() *AliexpressAscpFroQueryAPIRequest {
	return poolAliexpressAscpFroQueryAPIRequest.Get().(*AliexpressAscpFroQueryAPIRequest)
}

// ReleaseAliexpressAscpFroQueryAPIRequest 将 AliexpressAscpFroQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAscpFroQueryAPIRequest(v *AliexpressAscpFroQueryAPIRequest) {
	v.Reset()
	poolAliexpressAscpFroQueryAPIRequest.Put(v)
}
