package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscbusinessservicepricequeryAPIRequest 苹果查询服务供给报价 API请求
// alibaba.ssc.business.serviceprice.query
//
// 苹果查询服务供给报价
type AlibabasscbusinessservicepricequeryAPIRequest struct {
	model.Params
	// 查询参数
	_servicePriceQueryRequest *ServicePriceQueryRequest
}

// NewAlibabasscbusinessservicepricequeryRequest 初始化AlibabasscbusinessservicepricequeryAPIRequest对象
func NewAlibabasscbusinessservicepricequeryRequest() *AlibabasscbusinessservicepricequeryAPIRequest {
	return &AlibabasscbusinessservicepricequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscbusinessservicepricequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.business.serviceprice.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscbusinessservicepricequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscbusinessservicepricequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServicePriceQueryRequest is ServicePriceQueryRequest Setter
// 查询参数
func (r *AlibabasscbusinessservicepricequeryAPIRequest) SetServicePriceQueryRequest(_servicePriceQueryRequest *ServicePriceQueryRequest) error {
	r._servicePriceQueryRequest = _servicePriceQueryRequest
	r.Set("service_price_query_request", _servicePriceQueryRequest)
	return nil
}

// GetServicePriceQueryRequest ServicePriceQueryRequest Getter
func (r AlibabasscbusinessservicepricequeryAPIRequest) GetServicePriceQueryRequest() *ServicePriceQueryRequest {
	return r._servicePriceQueryRequest
}
