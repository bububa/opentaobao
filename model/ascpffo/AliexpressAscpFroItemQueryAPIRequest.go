package ascpffo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpFroItemQueryAPIRequest AliExpress销退单明细查询API API请求
// aliexpress.ascp.fro.item.query
//
// AE履约销退单明细查询API
type AliexpressAscpFroItemQueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentReverseOrderItemQuery *FulfillmentReverseOrderItemQueryDto
}

// NewAliexpressAscpFroItemQueryRequest 初始化AliexpressAscpFroItemQueryAPIRequest对象
func NewAliexpressAscpFroItemQueryRequest() *AliexpressAscpFroItemQueryAPIRequest {
	return &AliexpressAscpFroItemQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAscpFroItemQueryAPIRequest) Reset() {
	r._fulfillmentReverseOrderItemQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpFroItemQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.fro.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpFroItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpFroItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillmentReverseOrderItemQuery is FulfillmentReverseOrderItemQuery Setter
// dto
func (r *AliexpressAscpFroItemQueryAPIRequest) SetFulfillmentReverseOrderItemQuery(_fulfillmentReverseOrderItemQuery *FulfillmentReverseOrderItemQueryDto) error {
	r._fulfillmentReverseOrderItemQuery = _fulfillmentReverseOrderItemQuery
	r.Set("fulfillment_reverse_order_item_query", _fulfillmentReverseOrderItemQuery)
	return nil
}

// GetFulfillmentReverseOrderItemQuery FulfillmentReverseOrderItemQuery Getter
func (r AliexpressAscpFroItemQueryAPIRequest) GetFulfillmentReverseOrderItemQuery() *FulfillmentReverseOrderItemQueryDto {
	return r._fulfillmentReverseOrderItemQuery
}

var poolAliexpressAscpFroItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAscpFroItemQueryRequest()
	},
}

// GetAliexpressAscpFroItemQueryRequest 从 sync.Pool 获取 AliexpressAscpFroItemQueryAPIRequest
func GetAliexpressAscpFroItemQueryAPIRequest() *AliexpressAscpFroItemQueryAPIRequest {
	return poolAliexpressAscpFroItemQueryAPIRequest.Get().(*AliexpressAscpFroItemQueryAPIRequest)
}

// ReleaseAliexpressAscpFroItemQueryAPIRequest 将 AliexpressAscpFroItemQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAscpFroItemQueryAPIRequest(v *AliexpressAscpFroItemQueryAPIRequest) {
	v.Reset()
	poolAliexpressAscpFroItemQueryAPIRequest.Put(v)
}
