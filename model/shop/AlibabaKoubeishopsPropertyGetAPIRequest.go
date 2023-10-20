package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKoubeishopsPropertyGetAPIRequest 口碑店铺列表推荐 API请求
// alibaba.koubeishops.property.get
//
// 推荐用户附近的美食门店
type AlibabaKoubeishopsPropertyGetAPIRequest struct {
	model.Params
	// 入参
	_paramOpenApiSearchRequest *OpenApiSearchRequest
}

// NewAlibabaKoubeishopsPropertyGetRequest 初始化AlibabaKoubeishopsPropertyGetAPIRequest对象
func NewAlibabaKoubeishopsPropertyGetRequest() *AlibabaKoubeishopsPropertyGetAPIRequest {
	return &AlibabaKoubeishopsPropertyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaKoubeishopsPropertyGetAPIRequest) Reset() {
	r._paramOpenApiSearchRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKoubeishopsPropertyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.koubeishops.property.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKoubeishopsPropertyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaKoubeishopsPropertyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOpenApiSearchRequest is ParamOpenApiSearchRequest Setter
// 入参
func (r *AlibabaKoubeishopsPropertyGetAPIRequest) SetParamOpenApiSearchRequest(_paramOpenApiSearchRequest *OpenApiSearchRequest) error {
	r._paramOpenApiSearchRequest = _paramOpenApiSearchRequest
	r.Set("param_open_api_search_request", _paramOpenApiSearchRequest)
	return nil
}

// GetParamOpenApiSearchRequest ParamOpenApiSearchRequest Getter
func (r AlibabaKoubeishopsPropertyGetAPIRequest) GetParamOpenApiSearchRequest() *OpenApiSearchRequest {
	return r._paramOpenApiSearchRequest
}

var poolAlibabaKoubeishopsPropertyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaKoubeishopsPropertyGetRequest()
	},
}

// GetAlibabaKoubeishopsPropertyGetRequest 从 sync.Pool 获取 AlibabaKoubeishopsPropertyGetAPIRequest
func GetAlibabaKoubeishopsPropertyGetAPIRequest() *AlibabaKoubeishopsPropertyGetAPIRequest {
	return poolAlibabaKoubeishopsPropertyGetAPIRequest.Get().(*AlibabaKoubeishopsPropertyGetAPIRequest)
}

// ReleaseAlibabaKoubeishopsPropertyGetAPIRequest 将 AlibabaKoubeishopsPropertyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaKoubeishopsPropertyGetAPIRequest(v *AlibabaKoubeishopsPropertyGetAPIRequest) {
	v.Reset()
	poolAlibabaKoubeishopsPropertyGetAPIRequest.Put(v)
}
