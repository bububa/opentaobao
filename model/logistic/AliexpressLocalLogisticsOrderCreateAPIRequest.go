package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsOrderCreateAPIRequest create logistics order API请求
// aliexpress.local.logistics.order.create
//
// create logistics order
type AliexpressLocalLogisticsOrderCreateAPIRequest struct {
	model.Params
	// create logistics order's param
	_param1 *CreateOrderRequestTopDto
}

// NewAliexpressLocalLogisticsOrderCreateRequest 初始化AliexpressLocalLogisticsOrderCreateAPIRequest对象
func NewAliexpressLocalLogisticsOrderCreateRequest() *AliexpressLocalLogisticsOrderCreateAPIRequest {
	return &AliexpressLocalLogisticsOrderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressLocalLogisticsOrderCreateAPIRequest) Reset() {
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLocalLogisticsOrderCreateAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLocalLogisticsOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressLocalLogisticsOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// create logistics order&#39;s param
func (r *AliexpressLocalLogisticsOrderCreateAPIRequest) SetParam1(_param1 *CreateOrderRequestTopDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressLocalLogisticsOrderCreateAPIRequest) GetParam1() *CreateOrderRequestTopDto {
	return r._param1
}

var poolAliexpressLocalLogisticsOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressLocalLogisticsOrderCreateRequest()
	},
}

// GetAliexpressLocalLogisticsOrderCreateRequest 从 sync.Pool 获取 AliexpressLocalLogisticsOrderCreateAPIRequest
func GetAliexpressLocalLogisticsOrderCreateAPIRequest() *AliexpressLocalLogisticsOrderCreateAPIRequest {
	return poolAliexpressLocalLogisticsOrderCreateAPIRequest.Get().(*AliexpressLocalLogisticsOrderCreateAPIRequest)
}

// ReleaseAliexpressLocalLogisticsOrderCreateAPIRequest 将 AliexpressLocalLogisticsOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAliexpressLocalLogisticsOrderCreateAPIRequest(v *AliexpressLocalLogisticsOrderCreateAPIRequest) {
	v.Reset()
	poolAliexpressLocalLogisticsOrderCreateAPIRequest.Put(v)
}
