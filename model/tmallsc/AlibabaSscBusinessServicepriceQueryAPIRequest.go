package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscBusinessServicepriceQueryAPIRequest 苹果查询服务供给报价 API请求
// alibaba.ssc.business.serviceprice.query
//
// 苹果查询服务供给报价
type AlibabaSscBusinessServicepriceQueryAPIRequest struct {
	model.Params
	// 查询参数
	_servicePriceQueryRequest *ServicePriceQueryRequest
}

// NewAlibabaSscBusinessServicepriceQueryRequest 初始化AlibabaSscBusinessServicepriceQueryAPIRequest对象
func NewAlibabaSscBusinessServicepriceQueryRequest() *AlibabaSscBusinessServicepriceQueryAPIRequest {
	return &AlibabaSscBusinessServicepriceQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscBusinessServicepriceQueryAPIRequest) Reset() {
	r._servicePriceQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscBusinessServicepriceQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.business.serviceprice.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscBusinessServicepriceQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscBusinessServicepriceQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServicePriceQueryRequest is ServicePriceQueryRequest Setter
// 查询参数
func (r *AlibabaSscBusinessServicepriceQueryAPIRequest) SetServicePriceQueryRequest(_servicePriceQueryRequest *ServicePriceQueryRequest) error {
	r._servicePriceQueryRequest = _servicePriceQueryRequest
	r.Set("service_price_query_request", _servicePriceQueryRequest)
	return nil
}

// GetServicePriceQueryRequest ServicePriceQueryRequest Getter
func (r AlibabaSscBusinessServicepriceQueryAPIRequest) GetServicePriceQueryRequest() *ServicePriceQueryRequest {
	return r._servicePriceQueryRequest
}

var poolAlibabaSscBusinessServicepriceQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscBusinessServicepriceQueryRequest()
	},
}

// GetAlibabaSscBusinessServicepriceQueryRequest 从 sync.Pool 获取 AlibabaSscBusinessServicepriceQueryAPIRequest
func GetAlibabaSscBusinessServicepriceQueryAPIRequest() *AlibabaSscBusinessServicepriceQueryAPIRequest {
	return poolAlibabaSscBusinessServicepriceQueryAPIRequest.Get().(*AlibabaSscBusinessServicepriceQueryAPIRequest)
}

// ReleaseAlibabaSscBusinessServicepriceQueryAPIRequest 将 AlibabaSscBusinessServicepriceQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscBusinessServicepriceQueryAPIRequest(v *AlibabaSscBusinessServicepriceQueryAPIRequest) {
	v.Reset()
	poolAlibabaSscBusinessServicepriceQueryAPIRequest.Put(v)
}
