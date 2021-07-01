package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractIsvGatewayAPIRequest
isv调用gateway API请求
alibaba.interact.isv.gateway

isv能够调用jae本身的server */
type AlibabaInteractIsvGatewayAPIRequest struct {
	model.Params
}

// NewAlibabaInteractIsvGatewayRequest 初始化AlibabaInteractIsvGatewayAPIRequest对象
func NewAlibabaInteractIsvGatewayRequest() *AlibabaInteractIsvGatewayAPIRequest {
	return &AlibabaInteractIsvGatewayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvGatewayAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.isv.gateway"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvGatewayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
