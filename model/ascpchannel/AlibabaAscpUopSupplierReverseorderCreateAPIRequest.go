package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopsupplierreverseordercreateAPIRequest 商家ERP发起创建销退单服务 API请求
// alibaba.ascp.uop.supplier.reverseorder.create
//
// 商家在收到消费者实物退货后，在ERP发起创建销退单服务
type AlibabaascpuopsupplierreverseordercreateAPIRequest struct {
	model.Params
	// 逆向销退单创建请求
	_reverseCreateRequest *ReverseCreateRequest
}

// NewAlibabaascpuopsupplierreverseordercreateRequest 初始化AlibabaascpuopsupplierreverseordercreateAPIRequest对象
func NewAlibabaascpuopsupplierreverseordercreateRequest() *AlibabaascpuopsupplierreverseordercreateAPIRequest {
	return &AlibabaascpuopsupplierreverseordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuopsupplierreverseordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.reverseorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuopsupplierreverseordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuopsupplierreverseordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReverseCreateRequest is ReverseCreateRequest Setter
// 逆向销退单创建请求
func (r *AlibabaascpuopsupplierreverseordercreateAPIRequest) SetReverseCreateRequest(_reverseCreateRequest *ReverseCreateRequest) error {
	r._reverseCreateRequest = _reverseCreateRequest
	r.Set("reverse_create_request", _reverseCreateRequest)
	return nil
}

// GetReverseCreateRequest ReverseCreateRequest Getter
func (r AlibabaascpuopsupplierreverseordercreateAPIRequest) GetReverseCreateRequest() *ReverseCreateRequest {
	return r._reverseCreateRequest
}
