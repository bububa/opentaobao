package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsOrderInfoQueryAPIRequest query order details API请求
// aliexpress.local.logistics.order.info.query
//
// query order details
type AliexpressLocalLogisticsOrderInfoQueryAPIRequest struct {
	model.Params
	// param1
	_param1 *FindOrderRequestTopDto
}

// NewAliexpressLocalLogisticsOrderInfoQueryRequest 初始化AliexpressLocalLogisticsOrderInfoQueryAPIRequest对象
func NewAliexpressLocalLogisticsOrderInfoQueryRequest() *AliexpressLocalLogisticsOrderInfoQueryAPIRequest {
	return &AliexpressLocalLogisticsOrderInfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressLocalLogisticsOrderInfoQueryAPIRequest) Reset() {
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLocalLogisticsOrderInfoQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.order.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLocalLogisticsOrderInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressLocalLogisticsOrderInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// param1
func (r *AliexpressLocalLogisticsOrderInfoQueryAPIRequest) SetParam1(_param1 *FindOrderRequestTopDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressLocalLogisticsOrderInfoQueryAPIRequest) GetParam1() *FindOrderRequestTopDto {
	return r._param1
}

var poolAliexpressLocalLogisticsOrderInfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressLocalLogisticsOrderInfoQueryRequest()
	},
}

// GetAliexpressLocalLogisticsOrderInfoQueryRequest 从 sync.Pool 获取 AliexpressLocalLogisticsOrderInfoQueryAPIRequest
func GetAliexpressLocalLogisticsOrderInfoQueryAPIRequest() *AliexpressLocalLogisticsOrderInfoQueryAPIRequest {
	return poolAliexpressLocalLogisticsOrderInfoQueryAPIRequest.Get().(*AliexpressLocalLogisticsOrderInfoQueryAPIRequest)
}

// ReleaseAliexpressLocalLogisticsOrderInfoQueryAPIRequest 将 AliexpressLocalLogisticsOrderInfoQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressLocalLogisticsOrderInfoQueryAPIRequest(v *AliexpressLocalLogisticsOrderInfoQueryAPIRequest) {
	v.Reset()
	poolAliexpressLocalLogisticsOrderInfoQueryAPIRequest.Put(v)
}
